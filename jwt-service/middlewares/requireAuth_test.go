package middlewares

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequireAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	// add the cookie in the request
	cookie := &http.Cookie{
		Name:   "Authorization",
		Value:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc1MjA2MTksInN1YiI6MX0.5KfPKQT02C8jw8bY4p-OFbrdJPi1AOgZcqs-ibnumoI",
		MaxAge: 300,
	}
	c.Request.AddCookie(cookie)

	RequireAuth(c)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
