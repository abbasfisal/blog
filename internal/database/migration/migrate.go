package migration

import (
	"fmt"
	ArticleModel "github.com/abbasfisal/blog/internal/modules/article/models"
	UserModel "github.com/abbasfisal/blog/internal/modules/user/models"
	"github.com/abbasfisal/blog/pkg/database"
	"log"
)

func Migrate() {
	db := database.Db

	if err := db.AutoMigrate(&UserModel.User{}, &ArticleModel.Article{}); err != nil {
		log.Fatal("Migrate was failed: ", err)
		return
	}
	fmt.Println("migration done!")
}
