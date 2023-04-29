package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"runtime"
)

// https://blog.jetbrains.com/go/2021/04/30/how-to-use-docker-to-compile-go-from-goland/
func main() {
	jsoniter.ConfigCompatibleWithStandardLibrary.Marshal("3")
	fmt.Printf("Hello world from docker, os=%s, arch=%s\n\n", runtime.GOOS, runtime.GOARCH)
}
