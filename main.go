package main

import (
	"fmt"
	"github.com/ZakyHermawan/ginkaku/controller"
	"github.com/ZakyHermawan/ginkaku/repository"
	"github.com/ZakyHermawan/ginkaku/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	itemRepository = repository.NewItemRepository()
	itemService    = service.NewItemService(itemRepository)
	itemController = controller.NewItemController(itemService)
)

func main() {
	defer itemRepository.CloseDB()

	server := gin.New()
	server.Use(gin.Recovery())
	server.SetTrustedProxies([]string{"192.168.1.2"})

	server.GET("/orders", func(ctx *gin.Context) {
		allData, err := itemController.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, allData)
		}
	})
	server.POST("/orders", func(ctx *gin.Context) {
		if err := itemController.Save(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"message": "order created successfully"})
		}
	})
	server.PUT("/orders/:orderId", func(ctx *gin.Context) {
		if err := itemController.Update(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "order updated successfully"})
		}
	})
	server.DELETE("/orders/:orderId", func(ctx *gin.Context) {
		if err := itemController.Delete(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	})

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err.Error())
		return
	}

}
