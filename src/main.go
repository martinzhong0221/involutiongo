package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"involutiongo/src/service"
)

func main() {

	//DB, _ := sql.Open("mysql", "root:123456@tcp(39.108.118.123:3306)/test")

	dsn := "root:123456@tcp(39.108.118.123:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//product := entity.Product{ProductDesc: "健康险",ProductName: "Health Insurance"}
	//
	//result := db.Create(&product);
	//if result.Error!=nil{
	//	panic(result.Error)
	//}
	//fmt.Println(result.RowsAffected)
	//fmt.Println(product.Id)
	productService := service.NewProductService(db)
	product := productService.SelectById(1)
	fmt.Printf("%v", product)
}
