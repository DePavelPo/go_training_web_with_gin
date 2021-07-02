package article

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {

	articles := GetAllArticles()

	// Call the render function with the name of the template to render
	Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func GetArticle(c *gin.Context) {
	// Проверим валидность ID
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Проверим существование топика
		if article, err := GetArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			Render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// Если топика нет, прервём с ошибкой
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// При некорректном ID в URL, прервём с ошибкой
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
