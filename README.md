# Top Three Colours

project goal.... and in order to display a wide range of skills this project includes....

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

- go1.12.4

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

---

## Challenges

---

## Learnings

It's very fun to program in go and go-lint is fantastic for implementing some best practices. I enjoy the compiled nature to ensure I am shipping working code.

It is very worth taking time to study the Go standard library. For example, familiarity with `io.Writer` interface allows one to use `bytes.Buffer` in a test as `Writer`. Looking at the source code of a function like `fmt.Printf` helps to bridge this understanding. This project is a very good introduction and I definitely would further study the Go standard library.

---

## Some Helpful Resources

- [prominentcolor package to get top three colours](https://github.com/EdlinOrg/prominentcolor)
- [example implementation of prominentcolor](https://github.com/eddturtle/golangcode-site/blob/97c7260d42005ad05533ad3e188f6232869f20ed/content/posts/2019-07-04-find-common-colours-in-image.md)
- [testify to faciliate testing with assertions](https://github.com/stretchr/testify)
- [install submodules](https://stackoverflow.com/questions/35571079/go-how-to-import-package-from-github-and-build-without-go-get)
- [unit testing tips](https://www.red-gate.com/simple-talk/dotnet/software-testing/go-unit-tests-tips-from-the-trenches/#post-79965-_Toc519844651)
- [learn go with tests](https://github.com/quii/learn-go-with-tests/)
- [testify assertion reference](https://godoc.org/github.com/pgpst/pgpst/internal/github.com/stretchr/testify/assert)