package routes

import (
	"github.com/DePavelPo/go_training_web_with_gin/internal/article"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(c *gin.Engine) {
	g := c.Group("/v1/")

	// обработчик главного роута
	g.GET("/", article.ShowIndexPage)

	// Обработчик GET-запросов на /article/view/некоторый_article_id
	g.GET("/article/view/:article_id", article.GetArticle)
}
