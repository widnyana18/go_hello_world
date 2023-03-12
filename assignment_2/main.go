package main

import (
	"build_api/config"
	"build_api/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	init := config.ConnectDB()
	var db = &controller.Storage{DB: init}

	defer db.DB.Close()

	var router = gin.Default()

	router.GET("/orders", db.GetOrders)
	router.GET("/order/:id", db.GetOrderById)
	router.POST("/order/create", db.CreateOrder)
	router.PUT("/order/update", db.UpdateOrder)
	router.DELETE("/order/delete/:id", db.DeleteOrder)

	router.Run(":8080")
}
