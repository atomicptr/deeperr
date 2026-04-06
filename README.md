# deeperr

A lightweight Go error-wrapping library designed to provide stack traces and error codes while maintaining compatibility with the standard library

## Usage

```bash
$ go get atomicptr.dev/deeperr
```

```go
package main

import (
	"atomicptr.dev/deeperr"
)

// define your own error codes
const (
	ErrInternal deeperr.Code = 500
	ErrService  deeperr.Code = 503
	ErrAuth     deeperr.Code = 401
)

func main() {
	err := userController()
	deeperr.PrintStacktrace(err)
}

func userController() error {
	return deeperr.NewWithCode(ErrInternal, "controller: failed to resolve user", userService())
}

func userService() error {
	return deeperr.New("service: authentication layer failed", authProvider())
}

func authProvider() error {
	return deeperr.NewWithCode(ErrAuth, "provider: transient network failure", databaseLayer())
}

func databaseLayer() error {
	return deeperr.NewWithCode(ErrService, "database: connection pool exhausted", rootDriver())
}

func rootDriver() error {
	return deeperr.New("driver: socket hang up", nil)
}
```

this will now print:

```
E500 controller: failed to resolve user
    /home/christopher/dev/go/deeperr/example/example.go:20
service: authentication layer failed
    /home/christopher/dev/go/deeperr/example/example.go:24
E401 provider: transient network failure
    /home/christopher/dev/go/deeperr/example/example.go:28
E503 database: connection pool exhausted
    /home/christopher/dev/go/deeperr/example/example.go:32
driver: socket hang up
    /home/christopher/dev/go/deeperr/example/example.go:36
```

## License

MIT
