package middlewares_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
	"github.com/stretchr/testify/assert"
)

func TestPanicMiddleware(t *testing.T) {
	signature := ""
	router := gin.New()
	router.Use(gin.CustomRecovery(middlewares.RecoveryMiddleware))
	router.GET("/", func(c *gin.Context) {
		signature += "D"
		panic("abc")
	})
	router.GET("/test", func(c *gin.Context) {
		signature += "D"
		panic(middlewares.NewPanic(http.StatusAccepted, 1, "Test", errors.New("test")))
	})
	w := PerformRequest(router, "GET", "/", nil)
	assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError, "Expected Internal Server Error", w.Result().StatusCode)
	w = PerformRequest(router, "GET", "/test", nil)
	assert.Equal(t, w.Result().StatusCode, http.StatusAccepted, "Expected StatusAccepted", w.Result().StatusCode)
}
