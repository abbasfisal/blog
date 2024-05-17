package repositories

import ArticleModel "github.com/abbasfisal/blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
}
