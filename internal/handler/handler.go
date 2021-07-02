package handler

import (
	"github.com/DePavelPo/go_training_web_with_gin/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service
	gin     *gin.Engine
}

type service interface {
	GetAllArticles() []models.Article
	GetArticleByID(id int) (*models.Article, error)
}

func NewHandler(service service, gin *gin.Engine) {
	h := Handler{
		Service: service,
		gin:     gin,
	}

	g := h.gin.Group("/v1/")

	// обработчик главного роута
	g.GET("/", h.ShowIndexPage)

	// Обработчик GET-запросов на /article/view/некоторый_article_id
	g.GET("/article/view/:article_id", h.GetArticle)
}
