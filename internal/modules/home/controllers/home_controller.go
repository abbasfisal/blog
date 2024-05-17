package controllers

import (
	ArticleService "github.com/abbasfisal/blog/internal/modules/article/services"
	"github.com/abbasfisal/blog/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (c Controller) Index(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "home page",
		"featured": c.articleService.GetFeaturedArticles(),
		"stories":  c.articleService.GetStoriesArticle(),
	})

}
