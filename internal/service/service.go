package service

import "github.com/DePavelPo/go_training_web_with_gin/internal/models"

type repository interface{}

type Service interface {
	GetAllArticles() []models.Article
	GetArticleByID(id int) (*models.Article, error)
}

type service struct {
	repository repository
}

func NewService(repository repository) Service {
	return &service{
		repository: repository,
	}
}
