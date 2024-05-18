package routes

import (
	"github.com/abbasfisal/blog/internal/middlewares"
	controllers "github.com/abbasfisal/blog/internal/modules/article/controlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	ArticleController := controllers.New()

	r.GET("/articles/:id", ArticleController.Show)

	authGrp := r.Group("/articles")
	authGrp.Use(middlewares.IsAuth())
	{
		authGrp.GET("/create", ArticleController.Create)
		authGrp.POST("/create", ArticleController.Store)
	}
}
