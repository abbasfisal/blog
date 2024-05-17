package services

import (
	"github.com/abbasfisal/blog/internal/modules/user/reqeuests/auth"
	UserResponse "github.com/abbasfisal/blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(req auth.LoginRequest) (UserResponse.User, error)
}
