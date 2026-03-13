# Go run/test using Apple Containers

Quickly run and test your Linux-specific Go code from your Mac, using Apple Containers.

## Install

Note: You will need to install the [Apple Container CLI](https://github.com/apple/container) if you haven't already.

```sh
go install github.com/banksean/goac/cmd/...
```

## `gorunac`

`gorunac` works like `go run [...]`, but instead of building and executing a binary on your macOS host,
it cross-compiles a Linux binary and then executes the binary in an Apple Container.

## `gotestac`

`gotestac` works like `go test [...]`, but instead of building and executing tests on your macOS host,
it cross-compiles Linux test binaries and then executes the compiled tests in an Apple Container.


## Example usage

The [example](./example) directory contains some go code that conditionally compiles for Linux or macOS using `//go:build` tags.

```sh
> go run ./example # run it on macOS
Operating System: darwin
Running on macOS

> gorunac ./example # run it on Linux
Operating System: linux
Running on Linux
```

```sh
> go test ./example/... # test it on macOS
ok      github.com/banksean/goac/example        0.320s

> gotestac ./example/... # test it on Linux
PASS
```
