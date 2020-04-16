package test

import (
	"ModaLast/src/controllers/post"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllPosts(t *testing.T) {
	// Setup
	db := BaseTest()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := echo.HandlerFunc(post.GetAllPosts(db))

	// Assertions
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
