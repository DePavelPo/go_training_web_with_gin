package tests

import (
	"testing"

	"github.com/DePavelPo/go_training_web_with_gin/internal/models"
	"github.com/DePavelPo/go_training_web_with_gin/internal/repository"
	"github.com/DePavelPo/go_training_web_with_gin/internal/service"
)

func TestGetAllArticles(t *testing.T) {

	client := repository.NewClient()

	repository := repository.NewRepository(client)

	service := service.NewService(repository)

	t.Run("Success", func(t *testing.T) {
		var articleList = []models.Article{
			{ID: 1, Title: "Article 1", Content: "Article 1 body"},
			{ID: 2, Title: "Article 2", Content: "Article 2 body"},
		}

		retArticleList := service.GetAllArticles()

		if len(retArticleList) != len(articleList) {
			t.Fatalf("Success Test : Different sizes of arrays, must be the same")
		}

		for i, v := range retArticleList {
			if v.Content != articleList[i].Content ||
				v.ID != articleList[i].ID ||
				v.Title != articleList[i].Title {
				t.Fatalf("Success Test : Different sizes of arrays, must be the same")
				break
			}
		}
	})

	t.Run("ArrayLenFail", func(t *testing.T) {

		var articleList = []models.Article{
			{ID: 1, Title: "Article 1", Content: "Article 1 body"},
		}

		retArticleList := service.GetAllArticles()

		if len(retArticleList) == len(articleList) {
			t.Errorf("Array Len Fail Test : Same sizes of arrays, must be the different")
		}
	})

	t.Run("СontentFail", func(t *testing.T) {
		var articleList = []models.Article{
			{ID: 1, Title: "Article 1", Content: "Article 1 body"},
			{ID: 1, Title: "Article 2", Content: "Article 2 body"},
		}

		retArticleList := service.GetAllArticles()

		if len(retArticleList) != len(articleList) {
			t.Fatalf("Сontent Fail Test : Different sizes of arrays, must be the same")
		}
		var fail = true
		for i, v := range retArticleList {
			if v.Content != articleList[i].Content ||
				v.ID != articleList[i].ID ||
				v.Title != articleList[i].Title {
				fail = false
			}
		}
		if fail == true {
			t.Errorf("Сontent Fail Test : Array contents match, but must not match")
		}
	})

}
