package service

import (
	"gorm.io/gorm"
	"involutiongo/src/entity"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(p *gorm.DB) ProductService {
	return ProductService{db: p}
}

func (this *ProductService) SelectById(id uint) entity.Product {
	var product entity.Product
	this.db.First(&product, id)
	return product
}
