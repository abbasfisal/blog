package services

import (
	"errors"
	ArticleRepository "github.com/abbasfisal/blog/internal/modules/article/repositories"
	ArticleResponse "github.com/abbasfisal/blog/internal/modules/article/responses"
	"github.com/abbasfisal/blog/internal/modules/user/reqeuests/auth"
	UserResponse "github.com/abbasfisal/blog/internal/modules/user/responses"
)

type ArticleService struct {
	articleRepo ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepo: ArticleRepository.New(),
	}
}
func (a ArticleService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := a.articleRepo.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (a ArticleService) GetStoriesArticle() ArticleResponse.Articles {
	articles := a.articleRepo.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (a ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	article := a.articleRepo.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}
	return ArticleResponse.ToArticle(article), nil
}
