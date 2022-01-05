package main

import (
	"database/sql"
	"fmt"

	// _ 的意思是引入后面的包名而不直接使用这个包中定义的函数，变量等资源,
	//包在引入的时候会自动调用包的 init 函数以完成对包的初始化。
	//因此这里引入的数据库驱动包会自动去调用 init 函数，然后在 init 函数里面注册这个数据库驱动，
	//这样我们就可以在接下来的代码中直接使用这个数据库驱动了。
	_ "github.com/go-sql-driver/mysql"
)

func main_2() {
	/* sql.Open () 函数用来打开一个注册过的数据库驱动, mysql 表示数据库驱动名称，
	第二个参数是 DSN (Data Source Name)，它是 go-sql-driver 定义的一些数据库链接和配置信息
	支持格式如下:
	user@unix(/path/to/socket)/dbname?charset=utf8
	user:password@tcp(localhost:5555)/dbname?charset=utf8
	user:password@/dbname
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname */

	db, err := sql.Open("mysql", "root:Morso@2835@/mysite")
	checkErr(err)
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
