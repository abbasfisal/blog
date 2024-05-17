package seeder

import (
	"fmt"
	ArticleModel "github.com/abbasfisal/blog/internal/modules/article/models"
	UserModel "github.com/abbasfisal/blog/internal/modules/user/models"
	"github.com/abbasfisal/blog/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func Seed() {
	db := database.Connection()

	pass, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("password hashing was failed : ", err)
		return
	}
	user := UserModel.User{
		Model:    gorm.Model{},
		Name:     "ali",
		Email:    "ali@a.b",
		Password: string(pass),
	}
	db.Create(&user)

	fmt.Println("\t---[ user seeded successfully ]---")

	for i := 1; i <= 10; i++ {
		article := ArticleModel.Article{
			Model:   gorm.Model{},
			Title:   fmt.Sprintf("random title %d", i),
			Content: fmt.Sprintf("random content %d", i),
			UserID:  1,
		}

		db.Create(&article)
		log.Printf("article created successfully : [ title : %s , content : %s ]", article.Title, article.Content)
	}
	log.Println("seeder done ... ")
}
