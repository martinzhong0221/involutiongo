package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"involutiongo/src/entity"
)

func main() {

	DB, _ := sql.Open("mysql", "root:123456@tcp(39.108.118.123:3306)/test")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")

	var product entity.Product
	rows, e := DB.Query("select * from product where id =1")
	if e == nil {
		errors.New("query incur error")
	}
	for rows.Next() {
		e := rows.Scan(product.Id, product.ProductName, product.ProductDesc, product.UtcUpdate, product.UtcInsert)
		if e == nil {
			fmt.Println(json.Marshal(product))
		} else {
			fmt.Println(e)
		}
	}
	rows.Close()
	//单行查询操作
	//DB.QueryRow("select * from user where id=1").Scan(user.age, user.id, user.name, user.phone, user.sex)
}
