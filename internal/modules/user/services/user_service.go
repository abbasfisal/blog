package services

import (
	"errors"
	UserModel "github.com/abbasfisal/blog/internal/modules/user/models"
	"github.com/abbasfisal/blog/internal/modules/user/repositories"
	"github.com/abbasfisal/blog/internal/modules/user/reqeuests/auth"
	UserResponse "github.com/abbasfisal/blog/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepo: repositories.New(),
	}
}

func (u UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var userResponse UserResponse.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return userResponse, errors.New("error hashing the password")
	}
	var user = UserModel.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	newUser := u.userRepo.Create(user)
	if newUser.ID == 0 {
		return userResponse, errors.New("error on creating the user")
	}

	return UserResponse.ToUser(newUser), nil
}
