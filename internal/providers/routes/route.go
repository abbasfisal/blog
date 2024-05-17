package routes

import (
	ArticleRoutes "github.com/abbasfisal/blog/internal/modules/article/routes"
	HomeRoutes "github.com/abbasfisal/blog/internal/modules/home/routes"
	UserRoutes "github.com/abbasfisal/blog/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	HomeRoutes.Routes(r)
	ArticleRoutes.Routes(r)
	UserRoutes.Routes(r)
}
