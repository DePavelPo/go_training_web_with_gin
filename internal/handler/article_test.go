package handler

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DePavelPo/go_training_web_with_gin/internal/models"
	"github.com/DePavelPo/go_training_web_with_gin/internal/repository"
	"github.com/DePavelPo/go_training_web_with_gin/internal/service"
	//"github.com/DePavelPo/go_training_web_with_gin/internal/service"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {

	//client := repository.NewClient()

	//repository := repository.NewRepository(client)

	//service := service.NewService(repository)

	r := getRouter(true)

	NewHandler(service, r)

	//srv := Handler{Service: service}

	r.GET("/", ShowIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a GET request to the home page returns the list of articles
// in JSON format when the Accept header is set to application/json
func TestArticleListJSON(t *testing.T) {

	client := repository.NewClient()

	repository := repository.NewRepository(client)

	service := service.NewService(repository)

	r := getRouter(true)

	srv := Handler{Service: service}

	// Define the route similar to its definition in the routes file
	r.GET("/", srv.ShowIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is JSON which can be converted to
		// an array of Article structs
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []models.Article
		err = json.Unmarshal(p, &articles)

		return err == nil && len(articles) >= 2 && statusOK
	})
}

// Test that a GET request to an article page returns the article in XML
// format when the Accept header is set to application/xml
func TestArticleXML(t *testing.T) {

	client := repository.NewClient()

	repository := repository.NewRepository(client)

	service := service.NewService(repository)

	r := getRouter(true)

	srv := Handler{Service: service}

	// Define the route similar to its definition in the routes file
	r.GET("/article/view/:article_id", srv.GetArticle)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is JSON which can be converted to
		// an array of Article structs
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var a models.Article
		err = xml.Unmarshal(p, &a)

		return err == nil && a.ID == 1 && len(a.Title) >= 0 && statusOK
	})
}