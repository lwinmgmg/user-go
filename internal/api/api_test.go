package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/lwinmgmg/user-go/internal/api"
	"github.com/lwinmgmg/user-go/test"
)

var router *gin.Engine = nil

func init() {

	if router == nil {
		settings := test.GetTestEnv()
		router = api.SetupRouter(settings)
	}

}

func TestSetupRoute(t *testing.T) {
	settings := test.GetTestEnv()
	router := api.SetupRouter(settings)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}
