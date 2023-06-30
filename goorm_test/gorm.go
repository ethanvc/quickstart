package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Result struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func GormMain() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/test"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var result []Result
	db.Raw("DESCRIBE application_tab").Scan(&result)
	fmt.Printf("result = %v\n", result)
}
