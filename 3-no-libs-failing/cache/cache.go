package cache

// DataSource knows how to return the data from its original source when there is a cache miss.
type DataSource interface {
	Get(key string) string
}

// ReadThroughCache is a cache that uses the "read-through" pattern. When a client asks it for data,
// it will either return what it had stored previously or fetch, store, and return up to date data
// from its data source.
//
// This example is simplified in multiple ways:
//   - It can only store strings.
//   - It cannot GC old data. It will eventually fill up memory.
//   - It is not thread safe.
//   - It assumes the data source will always have the data in the event of a cache miss.
type ReadThroughCache struct {
	cachedData map[string]string
	dataSource DataSource
}

// NewReadThroughCache creates a new cache.
func NewReadThroughCache(dataSource DataSource) *ReadThroughCache {
	return &ReadThroughCache{
		cachedData: make(map[string]string),
		dataSource: dataSource,
	}
}

// Get looks for a value by key. If found, it returns it. If not found, it retrieves it from the
// data source and then returns what it retrieved.
func (c *ReadThroughCache) Get(key string) string {
	incorrectData := "y"

	// elem, found := c.cachedData[key]
	_, found := c.cachedData[key]
	if found {
		// Return the data that was found.
		// return elem
		return incorrectData
	}
	// Get the data from the data source and put it in the cached data before returning it.
	// data := c.dataSource.Get(key)
	data := incorrectData
	c.cachedData[key] = data
	return data
}
