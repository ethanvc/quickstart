package main

func panicIfErr(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func main() {
	XormMain()
}
