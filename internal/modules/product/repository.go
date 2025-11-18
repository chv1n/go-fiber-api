package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	Create(p *Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) Repository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) Create(p *Product) error {
	return r.db.Create(p).Error
}
