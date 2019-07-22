# app_sample

This is a simple go web server.

# Test

## 基准测试+功能测试+单元测试覆盖率
go test -v -cover -short -bench .
```
TEMP file :  /var/folders/j4/11rjs5s96cb_1dywr9bgc4xh0000gn/T/
2019/07/22 19:47:35 This is NEW version : 1.0.0
=== RUN   TestGetVersion
v1
--- PASS: TestGetVersion (0.00s)
=== RUN   TestLongTimeTestCase
--- SKIP: TestLongTimeTestCase (0.00s)
	main_test.go:30: Skip test when test in short mode.
=== RUN   TestParallel1
=== RUN   TestParallel2
--- PASS: TestParallel1 (1.00s)
	main_test.go:41: test parallel 1
--- PASS: TestParallel2 (1.00s)
	main_test.go:50: test parallel 2
goos: darwin
goarch: amd64
pkg: app_sample
BenchmarkGetVersion-4      	2000000000	         0.43 ns/op
BenchmarkUnmarshalJson-4   	  200000	      9842 ns/op
BenchmarkDecodeJson-4      	  200000	      8511 ns/op
PASS
coverage: 31.8% of statements
ok  	app_sample	5.808s
```

## 只运行基准测试

go test -run x -bench .
```
TEMP file :  /var/folders/j4/11rjs5s96cb_1dywr9bgc4xh0000gn/T/
2019/07/22 19:19:54 This is NEW version : 1.0.0
goos: darwin
goarch: amd64
pkg: app_sample
BenchmarkGetVersion-4   	2000000000	         0.34 ns/op
PASS
ok  	app_sample	0.731s

```


# httprouter

[httprouter](https://books.studygolang.com/advanced-go-programming-book/ch6-web/ch6-02-router.html)
[httprouter](https://www.kancloud.cn/kancloud/web-application-with-golang/44136)

```
https://books.studygolang.com/advanced-go-programming-book/ch6-web/ch6-02-router.html
https://www.kancloud.cn/kancloud/web-application-with-golang/44136
```
