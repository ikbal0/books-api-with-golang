package controllers

import (
	"fmt"
	"gin-api/database"
	"gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteCar godoc
// @Summary Delete car identified by given id
// @Description Delete the car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No content"
// @Router /cars/{Id} [delete]
func DeleteCar(c *gin.Context) {
	var db = database.GetDB()

	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// UpdateCar godoc
// @Summary Update car identified by given id
// @Description Update details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [patch]
func UpdateCar(c *gin.Context) {
	var db = database.GetDB()

	var car models.Car

	err := db.First(&car, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

// CreateCars godoc
// @Summary Post details for a given id
// @Description Get details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCars(c *gin.Context) {
	var db = database.GetDB()

	var input models.Car

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carInput := models.Car{Brand: input.Brand, Price: input.Price, CarType: input.CarType}
	db.Create(&carInput)

	c.JSON(http.StatusOK, gin.H{"data": carInput})
}

// GetOneCars godoc
// @Summary Get details for a given id
// @Description Get details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetOneCars(c *gin.Context) {
	var db = database.GetDB()

	var carOne []models.Car

	err := db.First(&carOne, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data one": carOne})
}

// GetAllCar godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(c *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car

	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car data:", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}
