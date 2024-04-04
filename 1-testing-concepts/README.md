# testing-concepts

Before learning how to use a mocking tool, it's good to understand what mocking is in the first place and when it's helpful to do mocking in (as opposed to other testing techniques).

This section of the tutorial is a basic introduction to Go software testing featuring some techniques and some tools.

Here's what this section isn't about:

* How to write and run a test. [Go's docs](https://go.dev/doc/tutorial/add-a-test) do this well already.
* How to generate and fill in a test table. [gotests](https://github.com/cweill/gotests) already does this well. This is the tool built into VS Code when you run "generate tests" from the command palette.

# Stubs, Mocks, and Martin Fowler

Software engineer Martin Fowler has a great article on his website, [Mocks Aren't Stubs](https://martinfowler.com/articles/mocksArentStubs.html).

The title of the article is about two particular tools that can be used for testing and the article uses Java as an example language. However, in addition to explaining what stubs and mocks are, and how to use them as tools when you write tests, there is content in the article that I think is universally useful, regardless of which language you're programming in and which testing tools you end up using.

In the article, Fowler differentiates these techniques from the mocking tools and calls the techniques "state verification" and "behavior verification".

As a Go programmer, I'd add one more technique to this - "input/output testing". You'll see why below.

## Input/output testing

Before getting into the other techniques, let's start with input/output testing, because this is the type of testing you're most likely to have already done with Go if you used only the built-in Go tooling.

When I say input/output testing, I'm talking about testing that a function returns a particular value given particular input. This works well with pure functions (ones that cause no side effects while they sort out which value to return), but it also works well in any scenario where you don't care how a function determines what to return, but that it returns the correct value.

In Go, we can use gotests to generate test tables to make it easy for us to quickly write many test cases for our functions and methods this way. In VS Code, we can invoke this tool via the command palette.

![image](https://github.com/mattwelke/go-mocking-tutorial/assets/7719209/09a90f43-52bf-4e7a-9779-e24eeb9041bb)

You end up with test tables that look like the ones featured in the [Go Wiki](https://go.dev/wiki/TableDrivenTests).

## State verification

So what about state verification then? If we aren't verifying the state returned by a function, what state are we verifying?

As Fowler explains in the article, this refers to the state of the entire system under test when the operation being tested finishes. In the article, he uses the example of an Order and a Warehouse, where an order is filled by a warehouse. The filling behavior is tested by checking that the warehouse has the inventory it should of a particular product after the order has been filled. He tests the Order and Warehouse classes at the same time. He doesn't use any [test doubles](https://martinfowler.com/bliki/TestDouble.html). For this testing technique, he doesn't need to. He explains in the article when you might decide it makes sense to use a test double.

Like in Fowler's example, with our Go testing, the system under test can be one thing tested in isolation or multiple things all tested at once. For us, this means Go structs. Later in the tutorial, you'll see a `ReadThroughCache` and `DataSource` as the two parts of our code, and it will be our goal to test `ReadThroughCache`.

## Behavior verification

This is the type of testing you do when the thing you care about is how one component in your system interacts with another. This is useful when we want to test more than just the "contract" of our code. For example, if we care more than just about what is returned from a function (input/output testing) or what state the system is in after the function finishes (state verification).

In the article, Fowler uses the example of a cache because a cache is an example of when you care about how things happen under the hood. In a system where the cache is a read-through cache (where the application isn't responsible for fetching data after cache misses) and the cache has a data source as a collaborator, you care about how many times it calls methods on that data source collaborator to get the data. When testing, you want to assert that it makes those calls only as often as it needs to (when the cache is empty or stale).

## Why mocking helps with behavior verification

"Mocking" refers to writing test doubles that are capable of recording their interactions with their collaborators. A mock is a stub in that it is another implementation of a collaborator of the system we want to test. Except, whereas a stub is defined as an alternate implementation that does what we tell it to for a particular test case, a mock is a stub where what we've told it to do is to record the interactions it had with our system under test.

## How mocking works in Go

In our case, that means a Go struct that tracks when and how its methods are called during a test case. Later in this tutorial, we will focus on this kind of testing, using a cache in Go as an example.

There are well-engineered, popular 3rd party libraries to do this for us too, so that we don't have to write it from scratch each time, which is error prone and introduces inconsistency in the codebase. Therefore, after showing how to write a mock data source from scratch, we will switch to using a popular Go mocking library to improve the tests.
