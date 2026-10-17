# Data structures and algorithms

<!-- progress:start -->[![progress](https://img.shields.io/badge/progress-0%2F282-lightgrey)](PROGRESS.md)<!-- progress:end -->

282 exercises in 33 chapters, for engineers preparing for coding interviews and contests. Every exercise is a stub that calls `panic("todo")` and has its own test. Write the code, make the test pass, move on.

The workbook explains each exercise and the answer key has a reference solution: [workbook](https://bloomcode.github.io/dsa/go/workbook.html) · [answer key](https://bloomcode.github.io/dsa/go/answer-key.html) · [concept guide](https://bloomcode.github.io/dsa/guide/).

## Start

1. Click **Use this template**, then **Create a new repository**. Use the template rather than forking: your copy gets its own history, can be private, and the progress workflow runs in it.
2. Clone your repository and install Go 1.24+.
3. Run `make test EX=1.1`. It fails until you write exercise 1.1.

## Work on an exercise

1. Open `chNN/<topic>.go` and find the block that starts with `// ==== EX 3.2`. Read its PROMPT, EXAMPLE and BIGO lines.
2. Replace the stub with your code. Keep the name and signature: the test calls them.
3. Run the test:

```sh
make test EX=3.2   # one exercise
make test EX=3     # one chapter
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
| Arrays, sets and searching | 1 Arrays and sets · 2 Ordered arrays · 3 Big O |
| Sorting and everyday Big O | 4 Bubble sort · 5 Selection sort · 6 Insertion sort · 7 Everyday Big O |
| Hash tables, stacks and queues | 8 Hash tables · 9 Stacks and queues |
| Recursion, dynamic programming and quicksort | 10 Recursion basics · 11 Writing recursive · 12 Dynamic programming · 13 Quicksort |
| Linked lists, trees, heaps and tries | 14 Linked lists · 15 Binary search trees · 16 Heaps · 17 Tries |
| Graphs, space and optimization | 18 Graphs · 19 Space complexity · 20 Optimization techniques |
| Sorting, speed and chance | 21 Mergesort · 22 Benchmarking · 23 Randomness |
| Caches and balanced trees | 24 Caching · 25 Red-black trees · 26 Treaps |
| Data bigger than memory | 27 B-trees · 28 External sort |
| Randomized algorithms | 29 Monte Carlo · 30 Hashing · 31 Rabin-Karp |
| Sets in a few bits | 32 Bit vectors · 33 Bloom filters |
