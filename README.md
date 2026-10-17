# Data structures and algorithms

<!-- progress:start -->[![progress](https://img.shields.io/badge/progress-0%2F282-lightgrey)](PROGRESS.md)<!-- progress:end -->

282 exercises in 26 chapters, for engineers preparing for coding interviews and contests. Every exercise is a stub that calls `panic("todo")` and has its own test. Write the code, make the test pass, move on.

The workbook explains each exercise and the answer key has a reference solution: [workbook](https://bloomcode.github.io/dsa/go/workbook.html) · [answer key](https://bloomcode.github.io/dsa/go/answer-key.html) · [concept guide](https://bloomcode.github.io/dsa/guide/).

## Start

1. Click **Use this template**, then **Create a new repository**. Use the template rather than forking: your copy gets its own history, can be private, and the progress workflow runs in it.
2. Clone your repository and install Go 1.24+.
3. Run `make test EX=11.1`. It fails until you write exercise 11.1.

## Work on an exercise

1. Open `chNN/<topic>.go` and find the block that starts with `// ==== EX 13.2`. Read its PROMPT, EXAMPLE and BIGO lines.
2. Replace the stub with your code. Keep the name and signature: the test calls them.
3. Run the test:

```sh
make test EX=13.2  # one exercise
make test EX=13    # one chapter
make test          # everything
make progress      # rewrite PROGRESS.md and the badge locally
```

Later exercises in a chapter often reuse earlier ones, so work through each chapter in order.

## Progress

Every push to `main` runs `.github/workflows/progress.yml`. It runs all the tests and commits [PROGRESS.md](PROGRESS.md), a tick per passing exercise, plus the badge at the top of this file. If its commit step fails with a permission error, open **Settings → Actions → General → Workflow permissions**, choose **Read and write permissions**, and push again.

## Getting updates

Fixes and new exercises land in [bloomcode/dsa-go](https://github.com/bloomcode/dsa-go). To pull them into your copy:

```sh
git remote add upstream https://github.com/bloomcode/dsa-go.git   # once
git fetch upstream
git merge upstream/main --allow-unrelated-histories   # the flag is only needed the first time
```

If a file you have worked on conflicts, keep your version of your code (`git checkout --ours <file>`) and take the upstream tests (`git checkout --theirs 'ch*/*_test.go'`), then `git add` and `git commit`.

## Layout

| Path | What |
| --- | --- |
| `chNN/<topic>.go` | Your code: one package per chapter, one block per exercise. |
| `chNN/<topic>_test.go` | One test per exercise, `TestEx<ch>_<num>`. Don't edit these. |
| `internal/check` | The tests' assertion helper. |
| `PROGRESS.md` | Written by the progress workflow: a tick per passing exercise. |
| `tools/progress.py` | Runs every test and writes PROGRESS.md (`make progress`). |
| `track.toml` | Chapter titles and parts. |

## Chapters

| Part | Chapters |
| --- | --- |
| Cost and measurement | 11 Counting steps · 12 Memory · 13 Timing |
| Sequences | 21 Arrays · 22 Sorted search · 23 Linked lists · 24 Stacks and queues |
| Sorting | 31 Simple sorts · 32 Divide and conquer |
| Hashing | 41 Hash tables · 42 Hash functions · 43 Text search |
| Trees | 51 Search trees · 52 Heaps · 53 Tries · 54 Balanced trees · 55 B-trees |
| Graphs | 61 Graph traversal · 62 Paths and order |
| Algorithm design | 71 Recursion · 72 Dynamic programming · 73 Greedy and transforms · 74 Randomized |
| Systems and scale | 81 Caching · 82 External sort · 83 Compact sets |
