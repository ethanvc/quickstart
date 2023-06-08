package main

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"testing"
)

// go test -bench=. -benchmem

type DataV0 struct {
	Name  string
	Age   int
	Class string
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

var v0 = &DataV0{
	Name:  "zhangsan",
	Age:   10,
	Class: "SanBan",
}

var v0msgpack = Must(msgpack.Marshal(v0))
var v0json = Must(json.Marshal(v0))

func BenchmarkMsgPackMarshalV0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(v0)
	}
}

func BenchmarkJsonMarshalV0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(v0)
	}
}

func BenchmarkMsgPackUnmarshalV0(b *testing.B) {
	var v DataV0
	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(v0msgpack, &v)
	}
}

func BenchmarkJsonUnmarshalV0(b *testing.B) {
	var v DataV0
	for i := 0; i < b.N; i++ {
		json.Unmarshal(v0json, &v)
	}
}
