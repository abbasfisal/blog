package middlewares

import (
	UserRepository "github.com/abbasfisal/blog/internal/modules/user/repositories"
	"github.com/abbasfisal/blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.New()
	return func(c *gin.Context) {
		authID := sessions.GET(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)
		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.Next()
	}
}
