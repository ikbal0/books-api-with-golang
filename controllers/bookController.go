package controllers

import (
	"books-api/database"
	"books-api/models"
	"fmt"
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
func DeleteBook(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var bookDelete models.Book
	err := db.First(&bookDelete, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	db.Delete(&bookDelete)

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
func UpdateBook(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var book []models.Book

	err := db.First(&book, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookUpdate := models.Book{NameBook: input.NameBook, Author: input.Author}
	db.Model(&input).Where("Id = ?", c.Param("id")).Updates(&bookUpdate)

	c.JSON(http.StatusOK, gin.H{"data": bookUpdate})
}

// CreateCars godoc
// @Summary Post details for a given id
// @Description Get details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Book body models.Book true "create car"
// @Success 200 {object} models.Book
// @Router /cars [post]
func CreateBook(c *gin.Context) {
	database.StartDB()

	var db = database.GetDB()

	var input models.Book

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookInput := models.Book{NameBook: input.NameBook, Author: input.Author}
	db.Create(&bookInput)

	c.JSON(http.StatusOK, gin.H{"data": bookInput})
}

// GetOneCars godoc
// @Summary Get details for a given id
// @Description Get details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Book
// @Router /cars/{Id} [get]
func GetOneBook(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var bookOne []models.Book

	err := db.First(&bookOne, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data one": bookOne})
}

// GetAllCar godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Book
// @Router /cars [get]
func GetAllBooks(c *gin.Context) {
	database.StartDB()
	var db = database.GetDB()

	var cars []models.Book

	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car data:", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}
