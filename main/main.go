package main

import (
	"fmt"
)

type IAbc interface {
	F1()
	F2()
}

type Abc struct {
}

func (Abc) F1() {
	fmt.Println("Abc.F1")
}

func (Abc) F2() {
	fmt.Println("Abc.F2")
}

type Bcd struct {
	IAbc
}

func main() {
	var bcd Bcd
	var abc Abc
	bcd.IAbc = abc
	bcd. = func(){
		fmt.Println("niew")
	}
	bcd.F2()
}
