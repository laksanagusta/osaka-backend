package order

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Save(order Order) (Order, error)
	Update(order Order) (Order, error)
	FindById(orderId int) (Order, error)
	BasketFindByOrderId(orderId int) ([]OrderProduct, error)
	SaveBasket(orderProduct OrderProduct) (OrderProduct, error)
	FindAll(page int, page_size int, q string) ([]Order, error)
	FindByStatus(status string, userId int) (Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(order Order) (Order, error) {
	err := r.db.Create(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) Update(order Order) (Order, error) {
	err := r.db.Save(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) FindById(orderId int) (Order, error) {
	order := Order{}
	err := r.db.First(&order, orderId).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) BasketFindByOrderId(orderId int) ([]OrderProduct, error) {
	orderProduct := []OrderProduct{}
	err := r.db.Preload("Order").Preload("Product").Where("order_id = ?", orderId).Find(&orderProduct).Error

	fmt.Println(orderProduct)

	if err != nil {
		return orderProduct, err
	}

	return orderProduct, nil
}

func (r *repository) SaveBasket(orderProduct OrderProduct) (OrderProduct, error) {
	err := r.db.Create(&orderProduct).Error

	if err != nil {
		return orderProduct, err
	}

	return orderProduct, nil
}

func (r *repository) FindAll(page int, pageSize int, q string) ([]Order, error) {
	var order []Order

	if q != "" {
		err := r.db.Where("order_number LIKE ?", "%"+q+"%").Offset(page).Limit(pageSize).Find(&order).Error
		if err != nil {
			return order, err
		}
	} else {
		err := r.db.Offset(page).Limit(pageSize).Find(&order).Error
		if err != nil {
			return order, err
		}
	}

	return order, nil
}

func (r *repository) FindByStatus(status string, userId int) (Order, error) {
	order := Order{}
	err := r.db.Where("status = ?", status).Where("user_id = ?", userId).First(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}
