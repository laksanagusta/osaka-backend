package order

import (
	"errors"
	"fmt"
)

type Service interface {
	Save(input OrderCreateInput) (Order, error)
	UpdateOrder(inputID FindById, inputData OrderCreateInput) (Order, error)
	FindById(orderId int) (Order, error)
	BasketFindByOrderId(orderId int) ([]OrderProduct, error)
	SaveBasket(input OrderCreateV2Input, orderId int) ([]OrderProduct, error)
	SaveBasketV2(input UpdateBasketInput, orderId int) ([]OrderProduct, error)
	FindAll(page int, page_size int, s string) ([]Order, error)
	FindByStatus(status string, userId int) (Order, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input OrderCreateInput) (Order, error) {
	order := Order{}
	order.Status = input.Status
	order.GrandTotal = input.GrandTotal
	order.CustomerName = input.CustomerName
	order.OrderNumber = input.OrderNumber

	newOrder, err := s.repository.Save(order)
	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func (s *service) UpdateOrder(inputID FindById, inputData OrderCreateInput) (Order, error) {
	order, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return order, err
	}

	order.Status = inputData.Status
	order.CustomerName = inputData.CustomerName
	order.GrandTotal = inputData.GrandTotal

	updatedOrder, err := s.repository.Update(order)
	if err != nil {
		return updatedOrder, err
	}

	return updatedOrder, nil

}

func (s *service) FindById(orderId int) (Order, error) {
	order, err := s.repository.FindById(orderId)
	if err != nil {
		return order, err
	}

	if order.ID == 0 {
		return order, errors.New("Order not found")
	}

	return order, nil
}

func (s *service) BasketFindByOrderId(orderId int) ([]OrderProduct, error) {
	orderProduct, err := s.repository.BasketFindByOrderId(orderId)
	if err != nil {
		return orderProduct, err
	}

	if len(orderProduct) == 0 {
		return orderProduct, errors.New("Order not found")
	}

	return orderProduct, nil
}

func (s *service) SaveBasket(input OrderCreateV2Input, orderId int) ([]OrderProduct, error) {
	basket := input.Basket

	orderProducts := []OrderProduct{}
	for _, basket := range basket {
		orderProduct := OrderProduct{}
		orderProduct.Qty = basket.Qty
		orderProduct.SubTotal = basket.SubTotal
		orderProduct.ProductID = basket.ProductID
		orderProduct.UnitPrice = basket.UnitPrice
		orderProduct.OrderID = orderId

		orderProduct, err := s.repository.SaveBasket(orderProduct)

		if err != nil {
			return orderProducts, err
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	fmt.Println(orderProducts)

	return orderProducts, nil
}

func (s *service) SaveBasketV2(input UpdateBasketInput, orderId int) ([]OrderProduct, error) {
	basket := input.Basket

	orderProducts := []OrderProduct{}
	for _, basket := range basket {
		orderProduct := OrderProduct{}
		orderProduct.Qty = basket.Qty
		orderProduct.SubTotal = basket.SubTotal
		orderProduct.ProductID = basket.ProductID
		orderProduct.UnitPrice = basket.UnitPrice
		orderProduct.OrderID = orderId

		orderProduct, err := s.repository.SaveBasket(orderProduct)

		if err != nil {
			return orderProducts, err
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	fmt.Println(orderProducts)

	return orderProducts, nil
}

func (s *service) FindAll(page int, pageSize int, q string) ([]Order, error) {
	order, err := s.repository.FindAll(page, pageSize, q)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (s *service) FindByStatus(status string, userId int) (Order, error) {
	order, err := s.repository.FindByStatus(status, userId)
	if err != nil {
		return order, err
	}

	if order.ID == 0 {
		return order, errors.New("Order not found")
	}

	return order, nil
}
