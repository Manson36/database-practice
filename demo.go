package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	//注意：程序在操作数据库的时候只需要用到database/sql，而不需要直接使用数据库驱动，所以程序在导入数据库驱动的时候将这个包的名字设置成下划线。
)

func main() {

	//"用户名：密码@[连接方式](主机名：端口号)/数据库名"
	//连接方式为tcp可以省略不写

	db, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/itcast")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	/*
	//操作一：执行数据操作语句
	sql := "insert into stu values (1,'tom')"
	result, _ := db.Exec(sql) //执行SQL语句
	n, _ := result.RowsAffected() //获取受影响的记录数

	fmt.Println("受影响的记录数", n)
	fmt.Printf("%T", result)
	*/

	/*
	//操作二：执行预处理
	stu := [2][2] string{{"3", "ketty"}, {"4", "rose"}}
	stmt, _ := db.Prepare("insert into stu values (?,?)") //获取预处理语句对象
	for _, s := range stu {
		stmt.Exec(s[0], s[1]) //调用预处理语句
	}
	*/

	/*
	//操作三：单行查询
	var id, name string
	row := db.QueryRow("select * from stu where id = 4")
	row.Scan(&id, &name)       //将row中的数据存到id，name中
	fmt.Println(id, "--", name)
	*/

	//操作四：多行查询
	rows, _ := db.Query("select * from stu")
	var id, name string
	for rows.Next() { //循环显示所有数据，rows.next()读取到数据返回true，否则返回false
		rows.Scan(&id, &name)
		fmt.Println(id, "--", name)
	}
}
