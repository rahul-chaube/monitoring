package routes

import (
	"github.com/rahul-chaube/monitoring/userService/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.POST("/device-token", controllers.StoreDeviceToken)
		userGroup.GET("/:id", controllers.GetUserByID)
		// userGroup.GET("/send-test-email", controllers.SendTestEmail)
		// userGroup.POST("/send-forwarding-email", controllers.SendForwardingEmail)
	}
}
