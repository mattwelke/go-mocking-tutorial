# no-libs

This directory shows setting up the code we want to test (the `ReadThroughCache` struct in `cache/cache.go`) and tests for it (in `cache/cache_test.go`).

This is the "no libraries" version. It's just plain Go to solve our problem of testing our code.

Our goal is to confirm that when our cache's `Get` method is called with a key corresponding to data it has, that it does not call `Get` on its data source. Also, that it does make this method call when it needs to because it doesn't yet have the data (but only once).

## Behavior verification

We want to use behavior verification as opposed to other techniques like state verification and input/output testing.

**Why not state verification?**

We're testing a cache, and caches don't normally track how often they return values to callers. This would use up memory needlessly as the program runs. We'd have to change the design of either our cache (to track how many times it calls `Get` on its data source) or our data source (so that it tracks how many times its `Get` method was called) just to facilitate testing. This would overcomplicate the code we intend to use in production.

**Why not input/output testing?**

This would be quick and easy code for us to write. We could use `gotests` to generate the test tables for us. But this would not address our use case. It would allow us to confirm that the cache returns the correct value given the correct key, but now *how* it returned it. It wouldn't allow us to confirm that it only needed to call `Get` on its data source when it didn't already have the data. From the outside, we wouldn't be able to see the difference between it handling one call to its `Get` method with the same key vs. two calls.

**Why behavior verification?**

It can be thought of as a kind of performance testing. Sometimes, when you code, you only care *what* code does when called with certain input. But sometimes, like in this case, you also care *how* it does that thing when called with that input.

Because the data source is a collaborator (i.e. dependency), we can use behavior verification for this. We measure the interactions between the cache and its data source - the calls to its `Get` method.
