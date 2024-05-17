package view

import (
	"github.com/abbasfisal/blog/internal/modules/user/helpers"
	"github.com/abbasfisal/blog/pkg/converters"
	"github.com/abbasfisal/blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("APP.Name")
	data["Errors"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["Olds"] = converters.StringToUrlValues(sessions.Flash(c, "olds"))
	data["Auth"] = helpers.Auth(c)
	return data
}
