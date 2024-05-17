package responses

import "github.com/abbasfisal/blog/internal/modules/user/models"

type User struct {
	ID    uint
	Image string
	Name  string
	Email string
}
type Users struct {
	Data []User
}

func ToUser(user models.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: "image",
	}
}
