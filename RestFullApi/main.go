package main

import (
	"restfullapi/config"
	"restfullapi/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	init := config.DBInit()
	var db = &controllers.DbCtrl{DB: init}

	var router = gin.Default()

	router.GET("/users", db.GetUsers)
	router.GET("/user/:id", db.GetUserById)
	router.POST("/create-user", db.CreateUser)
	router.PUT("/user/update", db.UpdateUser)
	router.DELETE("/user/delete/:id", db.DeleteUser)

	router.Run(":8080")
}
