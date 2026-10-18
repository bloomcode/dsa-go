"""Run every exercise's test and record which ones pass.

    python3 tools/progress.py          rewrite PROGRESS.md and the README badge
    python3 tools/progress.py --json   print the results only

Needs Python 3.11+ (tomllib). The progress workflow runs this on every push; you can run it
locally too. It never fails the build: an exercise that fails its test is simply not ticked.
"""
import datetime
import glob
import json
import os
import re
import subprocess
import sys
import tempfile
import tomllib
import xml.etree.ElementTree as ET
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
COMMENT = {"python": "#", "java": "//", "go": "//"}
BADGE_START, BADGE_END = "<!-- progress:start -->", "<!-- progress:end -->"


def chapter_key(ch):
    """Sort key for a chapter id: "2.5" -> (2, 5), 17 -> (17,)."""
    return tuple(int(x) for x in str(ch).split("."))


def ex_key(e):
    return chapter_key(e[0]), e[1]


def exercises(track):
    head = re.compile(rf"^\s*{re.escape(COMMENT[track['language']])} ==== EX (\d+(?:\.\d+)*)\.(\d+) \| (.+?) \| (\w+)")
    out = []
    for f in sorted(glob.glob(str(ROOT / track["sources"]))):
        if f.endswith("_test.go"):
            continue
        for line in Path(f).read_text().split("\n"):
            m = head.match(line)
            if m:
                out.append((m.group(1), int(m.group(2)), m.group(3), m.group(4)))
    return sorted(out, key=ex_key)


def run(cmd, timeout):
    try:
        r = subprocess.run(cmd, cwd=ROOT, capture_output=True, text=True, timeout=timeout)
        return r.stdout + r.stderr
    except subprocess.TimeoutExpired as e:
        return (e.stdout or b"").decode() if isinstance(e.stdout, bytes) else (e.stdout or "")


def results_python(ids):
    build_ok = True
    for f in glob.glob(str(ROOT / "solutions" / "*.py")):
        try:
            compile(Path(f).read_text(), f, "exec")
        except SyntaxError:
            build_ok = False
    with tempfile.TemporaryDirectory() as tmp:
        xml = Path(tmp) / "r.xml"
        run([sys.executable, "-m", "pytest", "-q", "-p", "no:cacheprovider", f"--junitxml={xml}"], 900)
        passed, failed = set(), set()
        if xml.exists():
            for case in ET.parse(xml).getroot().iter("testcase"):
                m = re.match(r"test_ex_((?:\d+_)*\d+)_(\d+)\b", case.get("name", ""))
                if not m:
                    continue
                eid = (m.group(1).replace("_", "."), int(m.group(2)))
                bad = any(child.tag in ("failure", "error") for child in case) or case.find("skipped") is not None
                (failed if bad else passed).add(eid)
        else:
            build_ok = False
    return build_ok, passed - failed


def results_java(ids):
    out = run(["sh", "run_tests.sh"], 900)
    passed = {(a, int(b)) for a, b in re.findall(r"^PASS (\d+(?:\.\d+)*)\.(\d+)$", out, re.M)}
    ran = passed | {(a, int(b)) for a, b in re.findall(r"^FAIL (\d+(?:\.\d+)*)\.(\d+)", out, re.M)}
    return bool(ran), passed


def results_go(ids, track):
    """One package per chapter. A panic stops the package's test binary, so rerun what didn't report."""
    build_ok, passed = True, set()
    pkg_of = {str(c["n"]): "./" + Path(c["file"]).parent.as_posix() for c in track["chapters"]}
    by_pkg = {}
    for ch, num in ids:
        by_pkg.setdefault(pkg_of[ch], []).append((ch, num))
    for pkg, todo in by_pkg.items():
        todo = set(todo)
        while todo:
            names = "|".join(f"TestEx{c.replace('.', '_')}_{n}" for c, n in sorted(todo, key=ex_key))
            out = run(["go", "test", "-count=1", "-json", "-run", f"^({names})$", pkg], 300)
            resolved, started, broken = set(), set(), False
            for line in out.splitlines():
                try:
                    ev = json.loads(line)
                except ValueError:
                    broken = broken or "build failed" in line or "setup failed" in line
                    continue
                if "[build failed]" in ev.get("Output", "") or "[setup failed]" in ev.get("Output", ""):
                    broken = True
                m = re.fullmatch(r"TestEx((?:\d+_)*\d+)_(\d+)", ev.get("Test", ""))
                if not m:
                    continue
                eid = (m.group(1).replace("_", "."), int(m.group(2)))
                if ev.get("Action") == "run":
                    started.add(eid)
                elif ev.get("Action") == "pass":
                    passed.add(eid)
                    resolved.add(eid)
                elif ev.get("Action") in ("fail", "skip"):
                    resolved.add(eid)
            resolved |= started - resolved  # a test that started but never finished crashed the binary
            if broken:
                build_ok = False
            if broken or not resolved:
                break
            todo -= resolved
    return build_ok, passed


def collect():
    track = tomllib.loads((ROOT / "track.toml").read_text())
    exs = exercises(track)
    ids = [(c, n) for c, n, _, _ in exs]
    lang = track["language"]
    build_ok, passed = results_go(ids, track) if lang == "go" else {"python": results_python, "java": results_java}[lang](ids)
    passed &= set(ids)
    return track, exs, build_ok, passed


def write(track, exs, build_ok, passed):
    total, done = len(exs), len(passed)
    today = datetime.date.today().isoformat()
    commit = os.environ.get("GITHUB_SHA", "")[:7]
    lines = ["# Progress", "", "<!-- Written by tools/progress.py on every push. Edits here are overwritten. -->", "",
             f"**{done} of {total} exercises pass ({100 * done // max(total, 1)}%).** "
             f"Last run {today}{f' at commit {commit}' if commit else ''}.", ""]
    if not build_ok:
        lines += ["> The code did not compile on the last run, so no test could pass. Run the tests locally to see the error.", ""]
    chapters = {str(c["n"]): c for c in track["chapters"]}
    for part in track["parts"]:
        lines += [f"## {part['name']}", "", "| Ch | Chapter | Done | Exercises |", "| ---: | --- | ---: | --- |"]
        for ch in map(str, part["chapters"]):
            mine = [(n, lvl) for c, n, _, lvl in exs if c == ch]
            got = sum((ch, n) in passed for n, _ in mine)
            ticks = " ".join(f"{'✅' if (ch, n) in passed else '⬜'}&nbsp;{ch}.{n}" for n, _ in mine)
            lines.append(f"| {ch} | {chapters[ch]['title']} | {got}/{len(mine)} | {ticks} |")
        lines.append("")
    (ROOT / "PROGRESS.md").write_text("\n".join(lines))
    readme = ROOT / "README.md"
    if readme.exists():
        colour = "2ea44f" if done == total else ("0969da" if done else "lightgrey")
        badge = f"[![progress](https://img.shields.io/badge/progress-{done}%2F{total}-{colour})](PROGRESS.md)"
        text = readme.read_text()
        new = re.sub(rf"{re.escape(BADGE_START)}.*?{re.escape(BADGE_END)}", f"{BADGE_START}{badge}{BADGE_END}", text, flags=re.S)
        if new != text:
            readme.write_text(new)


def main():
    track, exs, build_ok, passed = collect()
    if "--json" in sys.argv:
        ids = [(c, n) for c, n, _, _ in exs]
        print(json.dumps({"build_ok": build_ok, "total": len(ids), "passed": [f"{c}.{n}" for c, n in sorted(passed, key=ex_key)],
                          "failed": [f"{c}.{n}" for c, n in ids if (c, n) not in passed]}))
        return
    write(track, exs, build_ok, passed)
    print(f"{len(passed)} of {len(exs)} exercises pass" + ("" if build_ok else " (the code did not compile)"))


if __name__ == "__main__":
    main()
