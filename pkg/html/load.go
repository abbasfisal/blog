package html

import "github.com/gin-gonic/gin"

func LoadHtml(r *gin.Engine) {
	//internal/modules/moduleName/html/view.tmpl
	r.LoadHTMLGlob("internal/*/*/*/*html")
}
