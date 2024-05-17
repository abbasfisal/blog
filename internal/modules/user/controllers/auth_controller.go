package controllers

import (
	"github.com/abbasfisal/blog/internal/modules/user/reqeuests/auth"
	UserService "github.com/abbasfisal/blog/internal/modules/user/services"
	"github.com/abbasfisal/blog/pkg/converters"
	"github.com/abbasfisal/blog/pkg/errors"
	"github.com/abbasfisal/blog/pkg/html"
	"github.com/abbasfisal/blog/pkg/old"
	"github.com/abbasfisal/blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthController struct {
	userService UserService.UserServiceInterface
}

func New() *AuthController {
	return &AuthController{
		userService: UserService.New(),
	}
}
func (a AuthController) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
	return
}

func (a AuthController) HandleRegister(c *gin.Context) {
	//validate request
	var req auth.RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "olds", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}
	//create new user
	user, err := a.userService.Create(req)
	log.Println("user created successfully : ", user)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}
	c.Redirect(http.StatusFound, "/")
	return
}
