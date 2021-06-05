package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/docs/index.html

type DbOps struct {
	DB *gorm.DB
}

func init() {
}

func NewDbOps() *DbOps {
	dbOps := DbOps{}
	dsn := "root:123456@tcp(39.108.118.123:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbOps.DB = db
	fmt.Println("MySql init success")
	return &dbOps
}

func (p *DbOps) GetDB() *gorm.DB {
	return p.DB
}

//product := entity.Product{ProductDesc: "健康险",ProductName: "Health Insurance"}
//
//result := db.Create(&product);
//if result.Error!=nil{
//	panic(result.Error)
//}
//fmt.Println(result.RowsAffected)
//fmt.Println(product.Id)
//productService := service.NewProductService(db)
//product := productService.SelectById(1)
//fmt.Printf("%v", product)
