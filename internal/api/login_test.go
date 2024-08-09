package api_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func TestLoginApiSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "admin")
	req, _ := http.NewRequest("POST", "/user/api/v1/func/user/login", strings.NewReader(string(form.Encode())))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestLoginApiWrongFields(t *testing.T) {
	w := httptest.NewRecorder()
	form := url.Values{}
	form.Add("user", "admin")
	form.Add("pass", "admin")
	req, _ := http.NewRequest("POST", "/user/api/v1/func/user/login", strings.NewReader(string(form.Encode())))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 422, w.Code)
}

func TestLoginApiWrongPassword(t *testing.T) {
	w := httptest.NewRecorder()
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "wrongpassword")
	req, _ := http.NewRequest("POST", "/user/api/v1/func/user/login", strings.NewReader(string(form.Encode())))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
