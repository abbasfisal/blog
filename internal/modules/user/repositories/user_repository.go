package repositories

import (
	UserModel "github.com/abbasfisal/blog/internal/modules/user/models"
	"github.com/abbasfisal/blog/pkg/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		Db: database.Connection(),
	}
}

func (u UserRepository) Create(user UserModel.User) UserModel.User {
	var newUser UserModel.User

	u.Db.Create(&user).Scan(&newUser)
	return newUser
}
