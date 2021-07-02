package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DePavelPo/go_training_web_with_gin/internal/models"
	"github.com/DePavelPo/go_training_web_with_gin/internal/service"
	"github.com/gin-gonic/gin"
)

var tmpArticleList []models.Article

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("/home/pavelmuslimov/Desktop/GitHub-repos/go_training_web_with_gin/internal/app/templates/*")
	}
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
	tmpArticleList = service.ArticleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	service.ArticleList = tmpArticleList
}
