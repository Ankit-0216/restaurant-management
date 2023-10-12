package routes

import(
	"os"
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controller.GETUsers())
	incomingRoutes.GET("/users/:user_id", controller.GETUsers())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())

}