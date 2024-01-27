package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
	"github.com/stretchr/testify/assert"
)

func PerformRequest(r http.Handler, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestLoggerMiddleware(t *testing.T) {
	signature := ""
	router := gin.New()
	router.Use(middlewares.LoggerMiddleware)
	router.GET("/", func(c *gin.Context) {
		signature += "D"
	})
	w := PerformRequest(router, "GET", "/", nil)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK, "Status is not okay", w.Result().StatusCode)
}
