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

// 自定义表名
func (u *User) TableName() string {
	return getTableName(u.Id)
}

func init() {
	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册数据库连接
	orm.RegisterDataBase("instance1_db1", "mysql", "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4")
	orm.RegisterDataBase("instance1_db2", "mysql", "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4")
	orm.RegisterDataBase("instance2_db3", "mysql", "root:@tcp(127.0.0.1:3307)/?charset=utf8mb4")
	orm.RegisterDataBase("instance2_db4", "mysql", "root:@tcp(127.0.0.1:3307)/?charset=utf8mb4")

	// 注册模型
	orm.RegisterModel(new(User))
}

// 获取表名的函数，根据用户ID进行分表
func getTableName(userId int) string {
	tableIndex := userId % 2 // 两个表，取模2
	return fmt.Sprintf("user%d", tableIndex+1)
}

// 获取分库分表对应的数据库连接
func getAlias(userId, instanceId int) string {
	dbIndex := userId % 2 // 两个数据库，取模2
	return fmt.Sprintf("instance%d_db%d", instanceId, dbIndex+1)
}

func main() {
	userId := 2
	instanceId := 1                       // 使用第一个实例
	alias := getAlias(userId, instanceId) // 获取数据库连接别名

	o := orm.NewOrmUsingDB(alias)

	user := User{Id: userId, Name: "testuser"}
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println("Insert error:", err)
		return
	}

	var users []*User
	qs := o.QueryTable("user")
	_, err = qs.Filter("id", userId).All(&users)
	if err != nil {
		fmt.Println("Query error:", err)
		return
	}

	fmt.Println("Users:", users)
}
