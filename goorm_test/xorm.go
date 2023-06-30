package main

import "xorm.io/xorm"
import _ "github.com/go-sql-driver/mysql"

func XormMain() {
	engine, err := xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/test")
	panicIfErr(err)
	info, err := engine.DBMetas()
	panicIfErr(err)
	_ = info
}
