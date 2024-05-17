package helpers

import (
	UserRepository "github.com/abbasfisal/blog/internal/modules/user/repositories"
	UserResponse "github.com/abbasfisal/blog/internal/modules/user/responses"
	"github.com/abbasfisal/blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User

	authID := sessions.GET(c, "auth")
	userID, _ := strconv.Atoi(authID)

	var userRepo = UserRepository.New()

	user := userRepo.FindByID(userID)
	if user.ID == 0 {
		return response
	}

	return UserResponse.ToUser(user)
}
