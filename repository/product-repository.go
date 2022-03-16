package repository

import (
	"github.com/doffy007/price-comparison-api.git/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(p entities.Product) entities.Product
	UpdateProduct(p entities.Product) entities.Product
	DeleteProduct(p entities.Product)
	AllProduct() []entities.Product
	FindProductByTtitle(title string) entities.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(dbConc *gorm.DB) ProductRepository {
	return &productConnection{
		connection: dbConc,
	}
}

func (db *productConnection) InsertProduct(p entities.Product) entities.Product {
	db.connection.Save(&p)
	db.connection.Preload("Admin").Find(&p)
	return p
}

func (db *productConnection) UpdateProduct(p entities.Product) entities.Product {
	db.connection.Save(&p)
	db.connection.Preload("Admin").Find(&p)
	return p
}

func (db *productConnection) DeleteProduct(p entities.Product) {
	db.connection.Delete(&p)
}

func (db *productConnection) AllProduct() []entities.Product {
	var products []entities.Product
	db.connection.Preload("Admin").Find(&products)
	return products
}

func (db *productConnection) FindProductByTtitle(title string) entities.Product {
	var product entities.Product
	db.connection.Where("title = ?", title).Take(&product)
	return product
}