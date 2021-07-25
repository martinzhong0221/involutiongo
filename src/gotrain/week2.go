package gotrain

import (
	"database/sql"
	"errors"
	"involutiongo/src/entity"
)

// QueryName 当返回值是string int等数据类型时候，可以不处理异常。即调用者直接根据error来判空
func QueryName(db *sql.DB)(string,error) {
	var name string
	err := db.QueryRow("SELECT name from user",1).Scan(&name)
	if err != nil {
		if errors.Is(err,sql.ErrNoRows){
			println("找不到数据");
		} else {
			panic("发生其他错误")
		}
	}
	return name,err
}

// QueryProduct 当返回值是结构体的时候，可以把ErrNoRows吃掉，通过判断结构体为nil来体现没有找到数据。
func QueryProduct(db *sql.DB) *entity.Product {
	var product entity.Product
	err := db.QueryRow("SELECT * from products",1).Scan(&product)
	if err != nil {
		if errors.Is(err,sql.ErrNoRows){
			println("找不到数据");
			err = nil
		} else {
			panic("发生其他错误")
		}
	}
	return &product
}

