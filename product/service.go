package product

import (
	"errors"
)

type Service interface {
	Save(input ProductCreateInput) (Product, error)
	UpdateProduct(inputID FindById, inputData ProductCreateInput) (Product, error)
	FindById(productId int) (Product, error)
	SaveImage(productId int, fileLocation string) (Product, error)
	FindAll(page int, page_size int, s string) ([]Product, error)
	Delete(productId int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input ProductCreateInput) (Product, error) {
	product := Product{}
	product.Title = input.Title
	product.UnitPrice = input.UnitPrice
	product.Description = input.Description

	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) UpdateProduct(inputID FindById, inputData ProductCreateInput) (Product, error) {
	product, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return product, err
	}

	product.Title = inputData.Title
	product.Description = inputData.Description
	product.UnitPrice = inputData.UnitPrice

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil

}

func (s *service) FindById(productId int) (Product, error) {
	product, err := s.repository.FindById(productId)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("Product not found")
	}

	return product, nil
}

func (s *service) SaveImage(productId int, fileLocation string) (Product, error) {
	product, err := s.repository.FindById(productId)

	if err != nil {
		return product, err
	}

	product.Image = fileLocation

	updatedProduct, err := s.repository.Update(product)

	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *service) FindAll(page int, pageSize int, q string) ([]Product, error) {
	products, err := s.repository.FindAll(page, pageSize, q)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) Delete(productId int) (Product, error) {
	product, err := s.repository.Delete(productId)
	if err != nil {
		return product, err
	}

	return product, nil
}
