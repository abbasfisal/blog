package static

import "github.com/gin-gonic/gin"

func LoadStatic(r *gin.Engine) {
	r.Static("/assets", "./assets")
}
