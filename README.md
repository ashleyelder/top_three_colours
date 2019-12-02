# Top Three Colours

Project submission for [ code challenge](https://gist.github.com/ehmo/e736c827ca73d84581d812b3a27bb132).

## Sections
- [Design Choices](#design-choices)
- [Setting Up and Running Tests](#setting-up-and-running-tests)
- [Improvements](#improvements)
- [Challenges](#challenges)
- [Learnings](#learnings)
- [Some Helpful Resources](#some-helpful-resources)

---

## Design Choices

#### Tech Stack

- go1.12.4:
1. go supports concurrency instrinsic within the language
2. compiled language, it converts to machine code so it is faster and more efficient to execute than an interpreted language like python
3. nice and clean syntax unlike Java

#### Other Tools

- [prominentcolor](https://github.com/EdlinOrg/prominentcolor) (implementing algorithms and data structures is not the focus of this project)
- [testify](https://github.com/stretchr/testify)
- Go extension for Visual Studio Code that runs linter on save

---

## Setting Up and Running Tests

### Setup Project (requires installation of Docker)

- clone the repo
- cd into top_three_colours/
- run `docker build -t top_three_colours .`
- run `docker run -it top_three_colours`

### Run Tests

- **go test** shows just failure details plus a final result
- **go test -v** shows *all* test details plus a final result

#### Test Coverage

- **go test -cover** code coverage summary to the console
- **go test -coverprofile=cover.out** generate code coverage profile
- **go tool cover -html=cover.out** view graphical details in browser, would like a way to auto update this on save

---

## Improvements

Improve writing testable code and writing tests. Benchmark tests to assess performance.

Configurable number of workers as command line arguments using flags.Parse.

Breakout code into separate files for readability.

Graceful error handling, for instance when there is an issue like an empty line at the end of the input.txt file. Instead of returning `panic: runtime error: invalid memory address or nil pointer dereference` it could return a custom error message. For the sake of making my code more DRY I simply passed each default error message to the single checkError helper method. 

---

## Challenges

Synchronization

At the basic level, goroutines are easy to implement and have separate running activities. But to get more complicated functionality and having them line up to communicate is more challenging.

Also, tests are executed as goroutines as well.

---

## Learnings

It's very fun to program in go and I enjoy the compiled nature to ensure I am shipping working code.

It is very worth taking time to study the Go standard library. For example, familiarity with `io.Writer` interface allows one to use `bytes.Buffer` in a test as `Writer`. Looking at the source code of a function like `fmt.Printf` helps to bridge this understanding. This project is a very good introduction and I definitely would further study the Go standard library.

### Concurrency

Originally this program was sequential and had two main functions, one that did some computation and one that wrote some output and neither function calls the other. To make this a concurrent program, both functions are called and active at the same exact time. This models the real world and is faster.

---

## Some Helpful Resources

- [go concurrency beginner to advanced](https://github.com/golang/go/wiki/LearnConcurrency)
- [go concurrency patterns by Rob Pike](https://www.youtube.com/watch?v=f6kdp27TYZs&feature=youtu.be&t=617)