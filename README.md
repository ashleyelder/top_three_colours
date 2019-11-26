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

Improve writing testable code and writing tests.

---

## Challenges

---

## Learnings

It's very fun to program in go and go-lint is fantastic for implementing some best practices. I enjoy the compiled nature to ensure I am shipping working code.

It is very worth taking time to study the Go standard library. For example, familiarity with `io.Writer` interface allows one to use `bytes.Buffer` in a test as `Writer`. Looking at the source code of a function like `fmt.Printf` helps to bridge this understanding. This project is a very good introduction and I definitely would further study the Go standard library.

### Concurrency

goroutines are independent activities executing in a concurrenct Go program

channels allowe communication between goroutines

- ch <- x //  sends x to channel ch
- x = <-ch // x receives from ch
- <-ch // a receive statement; result is discarded
- close(ch) // no more values will be sent to this channel

<!-- Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop. -->

<!-- ^^^^ Channels should not be closed by their receiver because they don't know whether the sender is done sending. If the receiver closes a channel and the sender attempts to send again - the sender will panic. OTOH, if the sender closes the channel and receiver tries to read, the receiver reads nil.

Ergo, you close channels to let the receiver know there will be no more data. -->

originally this program was sequential and had two main functions, one that did some computation and one that wrote some output and neither function calls the other. to make this a concurrent program, calls to both functions are active at the same exact time.

---

## Some Helpful Resources

- [concurrency](https://github.com/golang/go/wiki/LearnConcurrency)
- [go concurrency patterns, channels @10:17](https://www.youtube.com/watch?v=f6kdp27TYZs&feature=youtu.be&t=617)
- [prominentcolor package to get top three colours](https://github.com/EdlinOrg/prominentcolor)
- [example implementation of prominentcolor](https://github.com/eddturtle/golangcode-site/blob/97c7260d42005ad05533ad3e188f6232869f20ed/content/posts/2019-07-04-find-common-colours-in-image.md)
- [testify to faciliate testing with assertions](https://github.com/stretchr/testify)
- [unit testing tips](https://www.red-gate.com/simple-talk/dotnet/software-testing/go-unit-tests-tips-from-the-trenches/#post-79965-_Toc519844651)
- [learn go with tests](https://github.com/quii/learn-go-with-tests/)
- [testify assertion reference](https://godoc.org/github.com/pgpst/pgpst/internal/github.com/stretchr/testify/assert)