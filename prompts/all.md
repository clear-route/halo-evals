Your job is to complete all of the coding exercises in this repository. The coding exercises are organized by programming language, and the directory structure reflects that:

- @/go/
- @/java/
- @/javascript/
- @/python/
- @/rust/

Within each of these directories is another set of directories, each representing a separate coding exercise that you need to complete. Each coding exercise has a file with an implementation stubbed, along with a test file (the tests should be failing initially).

To successfully complete an exercise you must fill out an implementation, run the tests, and all tests must pass. Do not alter the test file; it should be run as-is.

Do not use the "ask_followup_question" tool. Your job isn't done until the tests pass. Don't attempt completion until you run the tests and they pass.

You should start by reading the files in the `docs` directory within an exercises directory so that you understand the exercise, and then examine the stubbed out implementation file and test file.

To run the tests you can use the following commands:

go:
```sh
go test
```

java:
```sh
./gradlew test
```

javascript:
```sh
pnpm install --frozen-lockfile # If npm packages are not already installed.
pnpm test
```

python:
```sh
uv run python3 -m pytest -o markers=task [name]_test.py
```

rust:
```sh
cargo test
```
