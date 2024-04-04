# with-testify

This directory is the [`2-no-libs`](../2-no-libs/) directory, but changed to use the library [testify](https://github.com/stretchr/testify) for mocking. Instead of having to create the mocks manually, it makes function calls to the library to set up mock objects.

This has multiple advantages over the previous approach. They are as follows.

## Simpler mock struct state

This:

```go
// MockDataSource is a test double we can use with our tests, both simple and complex, to check
// whether all of our expectations were met.
type MockDataSource struct {
	// Tracks number of calls to Get with key as the map key.
	Calls map[string]int
	// The data to return for each key.
	ReturnData map[string]string
}
```

Turned into this:

```go
type MockDataSource struct {
	mock.Mock
}
```

This is because in the original testing approach, we had to come up with solutions to do all of the behavior verification we needed to do, across all of our test cases, and add things to the `MockDataSource` struct to enable this. For example, we had to add `Calls` in order to track the number of times certain functions were called. And we had to add `ReturnData` to control what data would be returned when those calls were made.

But in the new approach, the creators of testify have created building blocks in the form of methods on `mock.Mock` that we can call to set up our expectations and be able to assert based on them. Examples of using these methods are demonstrated below.

## Simpler measuring method calls

This:

```go
func (s *MockDataSource) Get(key string) string {
	if s.Calls == nil {
		s.Calls = make(map[string]int)
	}
	// Increment call count for the key
	s.Calls[key]++

	if s.ReturnData == nil {
		s.ReturnData = make(map[string]string)
	}
	// Return the data for the key.
	return s.ReturnData[key]
}
```

Turned into this:

```go
func (s *MockDataSource) Get(key string) string {
	args := s.Called(key)

	return args.String(0)
}
```

This is because testify gives us `Called` to help us track method calls. We don't have to add state to our mock object struct (like `Calls`) and mutate that state ourselves in order to track this.

## When writing each test case, simpler mock object setup

In the test case, this:

```go
// Create mock object, customized for this test case, capable of tracking what we need to know
// whether our expectations were met.
testDataSource := &MockDataSource{
	Calls: make(map[string]int),
	ReturnData: map[string]string{
		"a": "x",
	},
}
```

Turned into this:

```go
// Create the mock object that we will use to set up expectations on and assert that the
// expectations were met.
testDataSource := new(MockDataSource)
```

This is because with our mock object struct not having any state (because testify gives us the tools to set up and measure our expectations), there's no state to worry about setting up for the mock object in each test case.

## When writing each test case, simpler, standard assertions

When writing each test case, this:

```go
// Assert that the expectations were met.
expectedNumCalls := 1
if testDataSource.Calls["a"] != expectedNumCalls {
	t.Errorf("Expected Get to have been called %d time(s). It was called %d time(s).",
		expectedNumCalls, testDataSource.Calls["a"])
}
```

Turned into this:

```go
// Assert that the expectations were met.
testDataSource.AssertExpectations(t)
```

This is because testify tracked the expectations we set up. It knows how to determine whether the expectations were met when we call `AssertExpectations` at the end of the test case.
