package main

import(
	"os"
	"github.com/gin-gonic/gin"
	"restaurant-management/routes"
	"restaurant-management/database"
	"restaurant-management/models"
	"restaurant-management/middleware"
	"restaurant-management/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)


var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuItemRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
