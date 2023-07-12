package goormsqlstart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import "xorm.io/xorm"
import _ "github.com/go-sql-driver/mysql"

func TestXormStructField(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", MySqlDsn)
	panicIfErr(err)
	err = engine.Sync2(new(User))
	panicIfErr(err)
	_, err = engine.Insert(&User{EmbedField: EmbedField{
		UserName: "nihao",
	}})
	assert.NoError(t, err)
}

type User struct {
	UserId int64  `xorm:"bigint pk not null autoincr 'user_id'"`
	Name   string `xorm:"varchar(256) not null"`
	Addr   string `xorm:"varchar(256) not null"`
	Addr2  string `xorm:"varchar(256) not null"`
	// xorm will automatically convert struct/json when operating with db.
	EmbedField EmbedField `xorm:"varchar(512) json"`
}

type EmbedField struct {
	UserName string `json:"user_name"`
}
