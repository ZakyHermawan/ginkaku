package service

import (
	"github.com/ZakyHermawan/ginkaku/entity"
	"github.com/ZakyHermawan/ginkaku/repository"
)

type ItemService interface {
	CreateOrder(order *entity.Order) error
	UpdateOrder(order *entity.Order) error
	DeleteOrder(order *entity.Order) error
	GetAllOrder() ([]entity.Order, error)
}

type itemService struct {
	repository repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{
		repository: repo,
	}
}

func (service *itemService) CreateOrder(order *entity.Order) error {
	return service.repository.CreateOrder(order)
}

func (service *itemService) UpdateOrder(order *entity.Order) error {
	return service.repository.UpdateOrder(order)
}

func (service *itemService) DeleteOrder(order *entity.Order) error {
	return service.repository.DeleteOrder(order)
}

func (service *itemService) GetAllOrder() ([]entity.Order, error) {
	return service.repository.GetAllOrder()
}
