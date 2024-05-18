package services

import (
	"errors"
	ArticleModel "github.com/abbasfisal/blog/internal/modules/article/models"
	ArticleRepository "github.com/abbasfisal/blog/internal/modules/article/repositories"
	"github.com/abbasfisal/blog/internal/modules/article/requests"
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
func (a ArticleService) StoreAsUser(req requests.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	var article = ArticleModel.Article{
		Title:   req.Title,
		Content: req.Content,
		UserID:  user.ID,
	}
	newArticle := a.articleRepo.Create(article)
	if newArticle.ID == 0 {
		return response, errors.New("fail creating new article")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}
