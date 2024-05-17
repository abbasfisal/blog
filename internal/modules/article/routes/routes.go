package routes

import (
	controllers "github.com/abbasfisal/blog/internal/modules/article/controlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	ArticleController := controllers.New()

	r.GET("/articles/:id", ArticleController.Show)
}
