package main

func panicIfErr(err error) {
	if err == nil {
		return
	}
	panic(err)
}

// docker run --rm -it -e MYSQL_ALLOW_EMPTY_PASSWORD=true -p 127.0.0.1:10000:3306 mysql

const MySqlDsn = "root:@tcp(localhost:10000)/test"

func main() {
	XormMain()
	GormMain()
}
