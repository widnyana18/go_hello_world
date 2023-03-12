package controllers

import (
	"net/http"
	"restfullapi/model"

	"github.com/gin-gonic/gin"
)

func (gorm *DbCtrl) GetUsers(ctx *gin.Context) {
	var (
		persons []model.Person
		result  gin.H
	)

	gorm.DB.Find(&persons)

	if len(persons) <= 0 {
		result = gin.H{"message": "user not found", "count": 0}
	} else {
		result = gin.H{"message": persons, "count": len(persons)}
	}

	ctx.JSON(http.StatusOK, result)
}

func (gorm *DbCtrl) GetUserById(ctx *gin.Context) {
	var (
		person model.Person
		result gin.H
	)

	id := ctx.Param("id")
	err := gorm.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{"message": "user not found", "count": 0}
	} else {
		result = gin.H{"message": person, "count": 1}
	}

	ctx.JSON(http.StatusOK, result)
}

func (gorm *DbCtrl) CreateUser(ctx *gin.Context) {
	var (
		person model.Person
		result gin.H
	)

	fName := ctx.PostForm("firstName")
	lName := ctx.PostForm("lastName")
	person.FirstName = fName
	person.LastName = lName
	err := gorm.DB.Create(&person).Error

	if err != nil {
		result = gin.H{"message": "user already in use"}
	} else {
		result = gin.H{"message": person}
	}

	ctx.JSON(http.StatusOK, result)
}

func (gorm *DbCtrl) UpdateUser(ctx *gin.Context) {
	var (
		person    model.Person
		newPerson model.Person
		result    gin.H
	)

	id := ctx.Query("id")
	fName := ctx.PostForm("firstName")
	lName := ctx.PostForm("lastName")
	newPerson.FirstName = fName
	newPerson.LastName = lName

	err := gorm.DB.First(&person, id).Error

	if err != nil {
		result = gin.H{"message": "user not found"}
	}
	err = gorm.DB.Model(&person).Update(newPerson).Error

	if err != nil {
		result = gin.H{"message": "update failed"}
	} else {
		result = gin.H{"message": "update success", "data": person}
	}

	ctx.JSON(http.StatusOK, result)
}

func (gorm *DbCtrl) DeleteUser(ctx *gin.Context) {
	var (
		person model.Person
		result gin.H
	)

	id := ctx.Param("id")

	err := gorm.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{"message": "user not found"}
	}

	err = gorm.DB.Delete(&person).Error

	if err != nil {
		result = gin.H{"message": "delete failed"}
	} else {
		result = gin.H{"message": "delete success"}
	}

	ctx.JSON(http.StatusOK, result)
}
