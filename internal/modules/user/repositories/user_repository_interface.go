package repositories

import UserModel "github.com/abbasfisal/blog/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user UserModel.User) UserModel.User
}
