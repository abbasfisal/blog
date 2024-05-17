package routes

import (
	userCtrl "github.com/abbasfisal/blog/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	userController := userCtrl.New()

	r.GET("/register", userController.Register)
	r.POST("/register", userController.HandleRegister)
}
