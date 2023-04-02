package routers

import (
	"books-api/controllers"

	_ "books-api/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Car API
// @version 1.0
// @description This is a simple services for managing cars
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email devyad@gmail.com
// @license.name Apace 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	// Read
	router.GET("/books/:id", controllers.GetOneBook)
	// Read All
	router.GET("/books", controllers.GetAllBooks)
	// Create
	router.POST("/books", controllers.CreateBook)
	// Update
	router.PATCH("/books/:id", controllers.UpdateBook)
	// Delete
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
