package cache

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockDataSource struct {
	mock.Mock
}

func (s *MockDataSource) Get(key string) string {
	args := s.Called(key)

	return args.String(0)
}

func Test_CacheCallsDataSourceGet(t *testing.T) {
	// Create the mock object that we will use to set up expectations on and assert that the
	// expectations were met.
	testDataSource := new(MockDataSource)

	// Set up expectations.
	testDataSource.On("Get", "a").Return("x")

	// Call the code we are testing.
	sut := NewReadThroughCache(testDataSource)
	sut.Get("a")

	// Assert that the expectations were met.
	testDataSource.AssertExpectations(t)
}

func Test_CacheCallsDataSourceOnceForSameKey(t *testing.T) {
	// Create the mock object that we will use to set up expectations on and assert that the
	// expectations were met.
	testDataSource := new(MockDataSource)

	// Set up expectations.
	testDataSource.On("Get", "a").Return("x").Once()

	// Call the code we are testing.
	sut := NewReadThroughCache(testDataSource)
	sut.Get("a")
	sut.Get("a")

	// Assert that the expectations were met.
	testDataSource.AssertExpectations(t)
}
