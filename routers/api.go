package routers

import (
	"gin-api/controllers"

	_ "gin-api/docs"

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
	router.GET("/cars/:id", controllers.GetOneCars)
	// Read All
	router.GET("/cars", controllers.GetAllCars)
	// Create
	router.POST("/cars", controllers.CreateCars)
	// Update
	router.PATCH("/cars/:id", controllers.UpdateCar)
	// Delete
	router.DELETE("/cars/:id", controllers.DeleteCar)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
