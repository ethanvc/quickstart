package goormsqlstart

func panicIfErr(err error) {
	if err == nil {
		return
	}
	panic(err)
}

const MySqlDsn = "root:@tcp(localhost:10000)/test"
