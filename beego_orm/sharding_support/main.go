package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

// TableName 自定义表名
func (u *User) TableName() string {
	return fmt.Sprintf("`test_db_%d`.`test_tab`", u.Id%2)
}

func init() {
	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册数据库连接
	orm.RegisterDataBase("mysql0", "mysql", "root:@tcp(127.0.0.1:10000)/?charset=utf8mb4")
	orm.RegisterDataBase("mysql1", "mysql", "root:@tcp(127.0.0.1:10001)/?charset=utf8mb4")

	// 注册模型
	orm.RegisterModel(new(User))
}

// 获取分库分表对应的数据库连接
func getAlias(userId int) string {
	return fmt.Sprintf("mysql%d", userId%2)
}

func main() {
	orm.Debug = true
	u := &User{
		Id:   2,
		Name: "zhangsan",
	}
	o := orm.NewOrmUsingDB(getAlias(u.Id))

	_, err := o.Insert(u)
	if err != nil {
		fmt.Println("Insert error:", err)
		return
	}
	fmt.Printf("%d\n", u.Id)
}
