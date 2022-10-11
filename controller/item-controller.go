package controller

import (
	"github.com/ZakyHermawan/ginkaku/entity"
	"github.com/ZakyHermawan/ginkaku/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type ItemController interface {
	GetAll() ([]entity.Order, error)
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.ItemService
}

var validate *validator.Validate

func NewItemController(service service.ItemService) ItemController {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

func (c *controller) GetAll() ([]entity.Order, error) {
	return c.service.GetAllOrder()
}

func (c *controller) Save(ctx *gin.Context) error {

	var order entity.Order
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		return err
	}
	err = validate.Struct(order)
	if err != nil {
		return err
	}
	return c.service.CreateOrder(&order)
}

func (c *controller) Update(ctx *gin.Context) error {
	var order entity.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("orderId"), 0, 0)

	if err != nil {
		return err
	}
	order.ID = uint(id)

	err = validate.Struct(order)
	if err != nil {
		return err
	}

	err = c.service.UpdateOrder(&order)
	if err != nil {
		return err
	}
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var order entity.Order
	id, err := strconv.ParseUint(ctx.Param("orderId"), 0, 0)

	if err != nil {
		return err
	}
	order.ID = uint(id)
	return c.service.DeleteOrder(&order)
}
