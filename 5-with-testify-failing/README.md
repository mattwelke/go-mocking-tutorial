# with-testify-failing

This directory is the [`4-with-testify`](../4-with-testify/) directory, but with all tests removed except one and with the cache's `Get` implementation broken intentionally.

This is to show what one error message from a failing test looks like when testing is done this way.

This test output:

```
--- FAIL: Test_CacheCallsDataSourceOnceForSameKey (0.00s)
    /home/<user>/go-mocking-pres/3-no-libs-failing/cache/cache_test.go:48: Expected Get to have been called 1 time(s). It was called 0 time(s).
FAIL
FAIL	example.com/cache	0.001s
FAIL
```

Turns into this test output:

```
--- FAIL: Test_CacheCallsDataSourceOnceForSameKey (0.00s)
    /home/matt/go-mocking-pres/5-with-testify-failing/cache/cache_test.go:33: FAIL:	Get(string)
        		at: [/home/matt/go-mocking-pres/5-with-testify-failing/cache/cache_test.go:25]
    /home/matt/go-mocking-pres/5-with-testify-failing/cache/cache_test.go:33: FAIL: 0 out of 1 expectation(s) were met.
        	The code you are testing needs to make 1 more call(s).
        	at: [/home/matt/go-mocking-pres/5-with-testify-failing/cache/cache_test.go:33]
FAIL
FAIL	example.com/cache	0.002s
FAIL
```
