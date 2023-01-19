package product

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	FindById(productId int) (Product, error)
	FindByCode(code string) (Product, error)
	FindAll(page int, page_size int, q string) ([]Product, error)
	Delete(productId int) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindById(productId int) (Product, error) {
	product := Product{}
	err := r.db.First(&product, productId).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByCode(code string) (Product, error) {
	product := Product{}
	fmt.Println(code)
	err := r.db.Where("code = ?", code).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindAll(page int, pageSize int, q string) ([]Product, error) {
	var product []Product

	if q != "" {
		err := r.db.Offset(page).Limit(pageSize).Where("is_deleted = ?", 0).Where("title LIKE ?", "%"+q+"%").Find(&product).Error
		if err != nil {
			return product, err
		}
	} else {
		err := r.db.Offset(page).Limit(pageSize).Where("is_deleted = ?", 0).Find(&product).Error
		if err != nil {
			return product, err
		}
	}

	return product, nil
}

func (r *repository) Delete(productId int) (Product, error) {
	var product Product
	err := r.db.Model(&product).Where("id = ?", productId).Update("is_deleted", 1).Error

	if err != nil {
		return product, err
	}

	return product, nil
}
