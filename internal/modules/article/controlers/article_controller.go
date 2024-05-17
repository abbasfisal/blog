package controllers

import (
	ArticleService "github.com/abbasfisal/blog/internal/modules/article/services"
	"github.com/abbasfisal/blog/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (c Controller) Show(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		html.Render(ctx, http.StatusInternalServerError, "templates/errors/html/500", gin.H{
			"title": "server error",
			"error": "error converting id ",
		})
		return
	}

	article, err := c.articleService.Find(id)
	if err != nil {
		html.Render(ctx, http.StatusNotFound, "templates/errors/html/404", gin.H{
			"title": "page not found",
			"error": err.Error(),
		})
		return
	}
	html.Render(ctx, http.StatusOK, "modules/article/html/show",
		gin.H{
			"title":   "show article",
			"article": article,
		},
	)
	return
}

func (c Controller) Create(ctx *gin.Context) {

}
