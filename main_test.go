package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// Setup
	e := echo.New()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := hello(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello!\n", rec.Body.String())
	}
}
