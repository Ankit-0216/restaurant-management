package routes

import() {
	"github.com/gin-gonic/gin"
	controller"golang-restaurant-management/controllers"
}

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controller.GETFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GETFoods())
	incomingRoutes.POST("/users", controller.CreateFood())
	incomingRoutes.PATCH("/users/:user_id", controller.UpdateFood())
}