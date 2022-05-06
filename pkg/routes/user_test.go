package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

// does nothing
func TestGetUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	_ = c
}
