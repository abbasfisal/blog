package routes

import (
	"github.com/abbasfisal/blog/internal/middlewares"
	userCtrl "github.com/abbasfisal/blog/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	userController := userCtrl.New()

	guestGrp := r.Group("/")
	guestGrp.Use(middlewares.IsGuest())
	{
		r.GET("/register", userController.Register)
		r.POST("/register", userController.HandleRegister)

		r.GET("/login", userController.Login)
		r.POST("/login", userController.HandleLogin)
	}

	authGrp := r.Group("/")
	authGrp.Use(middlewares.IsAuth())
	{
		r.POST("/logout", userController.HandleLogout)
	}

}
