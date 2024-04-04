# go-mocking-tutorial

This is a tutorial/workshop on:

* How state verification and behavior verification work when testing code.
* Why mocks help with behavior verification.
* Why mocking libraries ([testify](https://github.com/stretchr/testify) used as an example) help with mocking while writing tests for Go code.

My inspiration for writing this tutorial comes from my experience first getting started with automated software testing early in my career, then getting familiar with how to do basic forms of it in Go (e.g. generating test tables for pure functions) while encountering insights throughout my career from industry leaders such as Martin Fowler (e.g. his [Mocks Aren't Stubs](https://martinfowler.com/articles/mocksArentStubs.html) article), and then learning more advanced forms of testing in Go.

## Getting started

Follow the tutorial by reading each README file in each directory in alphabetical order. Each README file explains testing concepts and describes the Go code contents of the directory.

If your system meets the following requirements, you can run the demonstration code yourself:

* Go 1.22+, or
* VS Code (with Go downloaded automatically by VS Code)

To run the tests, run `go test` in the subdirectory or click "run test" etc in VS Code next to the test you want to run.
