package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string `orm: "size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:Morso@2835@/mysite", 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "morso"}
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
}
