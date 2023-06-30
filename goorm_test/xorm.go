package main

import "xorm.io/xorm"
import _ "github.com/go-sql-driver/mysql"

type User struct {
	UserId int64  `xorm:"bigint not null autoincr 'user_id'"`
	Name   string `xorm:"varchar(256) not null"`
	Addr   string `xorm:"varchar(256) not null"`
	Addr2  string `xorm:"varchar(256) not null"`
}

func XormMain() {
	engine, err := xorm.NewEngine("mysql", MySqlDsn)
	panicIfErr(err)
	info, err := engine.DBMetas()
	_ = info
	panicIfErr(err)

	err = engine.Sync2(new(User))
	panicIfErr(err)
}
