package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/repositories"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/ping")
	c := &controller{e, repositories.MockUsersRepository()}

	// Assertions
	if assert.NoError(t, c.Ping(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Pong!", rec.Body.String())
	}
}
