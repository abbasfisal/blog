package services

import (
	ArticleResponse "github.com/abbasfisal/blog/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticle() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
}
