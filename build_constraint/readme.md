build constraint有两种约束方式：

1. 文件名去除`_test.go`后缀后，后缀带有`*_GOOS`, `*_GOARCH`, `*_GOOS_GOARCH`，example: `source_windows_amd64.go`
2. 带有 `//go:build` 注释行，也就是build tag的文件

