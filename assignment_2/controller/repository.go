package controller

import (
	"build_api/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (src *Storage) GetOrders(ctx *gin.Context) {
	var (
		orders []model.Orders
		result gin.H
	)

	src.DB.Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{"message": "Order not found", "count": 0}
	} else {
		result = gin.H{"data": orders, "count": len(orders)}
	}

	ctx.JSON(http.StatusOK, result)
}

func (src *Storage) GetOrderById(ctx *gin.Context) {
	var (
		order  model.Orders
		result gin.H
	)

	orderId := ctx.Param("orderId")
	err := src.DB.Where("OrderId = ?", orderId).First(&order).Error

	if err != nil {
		result = gin.H{"message": "Order not found", "count": 0}
	} else {
		result = gin.H{"data": order, "count": 1}
	}

	ctx.JSON(http.StatusOK, result)
}

func (src *Storage) CreateOrder(ctx *gin.Context) {
	var result gin.H

	custName := ctx.PostForm("customerName")
	itemCode := ctx.PostForm("itemCode")
	desc := ctx.PostForm("description")
	qty, err := strconv.Atoi(ctx.PostForm("quantity"))
	if err != nil {
		log.Println(err.Error())
	}

	var items = []model.Items{
		{ItemCode: itemCode, Description: desc, Quantity: qty},
	}
	var order = model.Orders{CustomerName: custName, OrderedAt: time.Now(), ItemList: items}

	err = src.DB.Create(&order).Error

	if err != nil {
		result = gin.H{"message": "Order already in use"}
	} else {
		result = gin.H{"message": "Create success", "data": order}
	}

	ctx.JSON(http.StatusOK, result)
}

func (src *Storage) UpdateOrder(ctx *gin.Context) {
	var (
		order  model.Orders
		result gin.H
	)

	orderId := ctx.Query("orderId")
	custName := ctx.PostForm("customerName")
	itemCode := ctx.PostForm("itemCode")
	desc := ctx.PostForm("description")
	qty, err := strconv.Atoi(ctx.PostForm("quantity"))
	if err != nil {
		log.Println(err.Error())
	}

	var items = []model.Items{
		{ItemCode: itemCode, Description: desc, Quantity: qty},
	}
	var newOrder = model.Orders{CustomerName: custName, OrderedAt: time.Now(), ItemList: items}

	err = src.DB.First(&order, orderId).Error

	if err != nil {
		result = gin.H{"message": "Order not found"}
	}
	err = src.DB.Model(&order).Update(newOrder).Error

	if err != nil {
		result = gin.H{"message": "update failed"}
	} else {
		result = gin.H{"message": "update success", "data": order}
	}

	ctx.JSON(http.StatusOK, result)
}

func (src *Storage) DeleteOrder(ctx *gin.Context) {
	var (
		order  model.Orders
		result gin.H
	)

	orderId := ctx.Param("orderId")

	err := src.DB.First(&order, orderId).Error
	if err != nil {
		result = gin.H{"message": "Order not found"}
	}

	err = src.DB.Delete(&order).Error

	if err != nil {
		result = gin.H{"message": "delete failed"}
	} else {
		result = gin.H{"message": "delete success"}
	}

	ctx.JSON(http.StatusOK, result)
}
