# simple-cobra-app [![Build Status](https://travis-ci.org/testeradjust/simple-cobra-app.svg?branch=master)](https://travis-ci.org/testeradjust/simple-cobra-app)

For Travis CI testing.

Skip travis build by providing `[skip travis]` in the commit message, like so:
```
[skip travis] README Update
```

Services/processes successfully integrated so far:

- aerospike
- redis
- kafka
- postgres
- golang CI
- easyjson

Sometimes, after Travis CI finishes the build cycle, it seems to fail to report the status to GitHub, as described here:
https://github.com/travis-ci/travis-ci/issues/10204
