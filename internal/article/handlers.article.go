package article

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {

	articles := GetAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(

		// Set the HTTP status to 200 (OK)
		http.StatusOK,

		// Use the index.html template
		"index.html",

		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func GetArticle(c *gin.Context) {
	// Проверим валидность ID
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Проверим существование топика
		if article, err := GetArticleByID(articleID); err == nil {
			// Вызовем метод HTML из Контекста Gin для обработки шаблона
			c.HTML(
				// Зададим HTTP статус 200 (OK)
				http.StatusOK,
				// Используем шаблон index.html
				"article.html",
				// Передадим данные в шаблон
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)

		} else {
			// Если топика нет, прервём с ошибкой
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// При некорректном ID в URL, прервём с ошибкой
		c.AbortWithStatus(http.StatusNotFound)
	}
}
