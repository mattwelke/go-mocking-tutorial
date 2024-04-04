package cache

import (
	"testing"
)

// MockDataSource is a test double we can use with our tests, both simple and complex, to check
// whether all of our expectations were met.
type MockDataSource struct {
	// Tracks number of calls to Get with key as the map key.
	Calls map[string]int
	// The data to return for each key.
	ReturnData map[string]string
}

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

func Test_CacheCallsDataSourceGet(t *testing.T) {
	// Create mock object, customized for this test case, capable of tracking what we need to know
	// whether our expectations were met.
	testDataSource := &MockDataSource{
		Calls: make(map[string]int),
		ReturnData: map[string]string{
			"a": "x",
		},
	}

	// Call the code we are testing.
	sut := NewReadThroughCache(testDataSource)
	sut.Get("a")

	// Assert that the expectations were met.
	if testDataSource.Calls["a"] < 1 {
		t.Error("Expected Get to have been called at least once. It was never called.")
	}
}

func Test_CacheCallsDataSourceOnceForSameKey(t *testing.T) {
	// Create mock object, customized for this test case, capable of tracking what we need to know
	// whether our expectations were met.
	testDataSource := &MockDataSource{
		Calls: make(map[string]int),
		ReturnData: map[string]string{
			"a": "x",
		},
	}

	// Call the code we are testing.
	sut := NewReadThroughCache(testDataSource)
	sut.Get("a")
	sut.Get("a")

	// Assert that the expectations were met.
	expectedNumCalls := 1
	if testDataSource.Calls["a"] != expectedNumCalls {
		t.Errorf("Expected Get to have been called %d time(s). It was called %d time(s).",
			expectedNumCalls, testDataSource.Calls["a"])
	}
}
