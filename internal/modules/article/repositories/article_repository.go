package repositories

import (
	ArticleModel "github.com/abbasfisal/blog/internal/modules/article/models"
	"github.com/abbasfisal/blog/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	Db *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		Db: database.Connection(),
	}
}

func (a ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article
	a.Db.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

func (a ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article

	a.Db.Joins("User").First(&article, id)
	return article
}

func (a ArticleRepository) Create(article ArticleModel.Article) ArticleModel.Article {
	var newArticle ArticleModel.Article

	a.Db.Create(&article).Scan(&newArticle)
	return newArticle
}
