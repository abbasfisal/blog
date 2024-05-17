package responses

import (
	"fmt"
	"github.com/abbasfisal/blog/internal/modules/article/models"
	UserResponse "github.com/abbasfisal/blog/internal/modules/user/responses"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      UserResponse.User
}
type Articles struct {
	Data []Article
}

func ToArticle(article models.Article) Article {
	return Article{
		ID:        article.ID,
		Image:     "/assets/img/default-img.jpg",
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: fmt.Sprintf("%d/%d/%d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      UserResponse.ToUser(article.User),
	}
}

func ToArticles(articles []models.Article) Articles {
	var response Articles
	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}
	return response
}
