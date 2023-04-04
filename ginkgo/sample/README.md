1. Running `ginkgo -p -vv --fail-fast`, you will see `SynchronizedAfterSuite` are interrupted

```bash
Summarizing 15 Failures:
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[FAIL] Framework E2E [It] should run a failed test
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:24
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11
[INTERRUPTED] [SynchronizedAfterSuite] 
/home/weizhichen/go-playground/ginkgo/sample/sample_test.go:11

Ran 1 of 2 Specs in 6.035 seconds
FAIL! - Interrupted by Other Ginkgo Process -- 0 Passed | 1 Failed | 0 Pending | 1 Skipped
```

2. Comment out SynchronizedAfterSuite, and run `ginkgo -p -vv --fail-fast` again, you will see that the Serial spec is interrupted this time: 

```bash
------------------------------

Summarizing 2 Failures:
  [FAIL] Framework E2E [It] should run a failed test
  /home/weizhichen/go-playground/ginkgo/sample/sample_test.go:24
  [INTERRUPTED] Framework E2E [It] should run serial test
  /home/weizhichen/go-playground/ginkgo/sample/sample_test.go:27

Ran 2 of 2 Specs in 2.056 seconds
FAIL! - Interrupted by Other Ginkgo Process -- 0 Passed | 2 Failed | 0 Pending | 0 Skipped
```