package services

import (
	"github.com/abbasfisal/blog/internal/modules/article/requests"
	ArticleResponse "github.com/abbasfisal/blog/internal/modules/article/responses"
	UserResponse "github.com/abbasfisal/blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticle() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(req requests.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}
