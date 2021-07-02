package service

import (
	"errors"

	"github.com/DePavelPo/go_training_web_with_gin/internal/models"
)

var ArticleList = []models.Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func (s *service) GetAllArticles() []models.Article {

	return ArticleList
}

func (s *service) GetArticleByID(id int) (*models.Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
