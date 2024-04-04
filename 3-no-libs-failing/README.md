# no-libs-failing

This directory is the [`2-no-libs`](../2-no-libs/) directory, but with all tests removed except one and with the cache's `Get` implementation broken intentionally.

This is to show what one error message from a failing test looks like when testing is done this way.

The test output:

```
--- FAIL: Test_CacheCallsDataSourceOnceForSameKey (0.00s)
    /home/<user>/go-mocking-pres/3-no-libs-failing/cache/cache_test.go:48: Expected Get to have been called 1 time(s). It was called 0 time(s).
FAIL
FAIL	example.com/cache	0.001s
FAIL
```
