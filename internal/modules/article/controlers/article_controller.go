package controllers

import (
	"fmt"
	"github.com/abbasfisal/blog/internal/modules/article/requests"
	ArticleService "github.com/abbasfisal/blog/internal/modules/article/services"
	"github.com/abbasfisal/blog/internal/modules/user/helpers"
	"github.com/abbasfisal/blog/pkg/converters"
	"github.com/abbasfisal/blog/pkg/errors"
	"github.com/abbasfisal/blog/pkg/html"
	"github.com/abbasfisal/blog/pkg/old"
	"github.com/abbasfisal/blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"log"
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
	html.Render(ctx, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "create article",
	})
	return
}

func (c Controller) Store(ctx *gin.Context) {
	//validate request
	var req requests.StoreRequest
	if err := ctx.ShouldBind(&req); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "olds", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(ctx)

	//create new article
	article, err := c.articleService.StoreAsUser(req, user)
	log.Println("article created successfully : ", user)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/articles/create")
		return
	}

	ctx.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
	return

}
