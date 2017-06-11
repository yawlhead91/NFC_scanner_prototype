package apis

import (
	"testing"
	"net/http"
)

func TestAuth(t *testing.T) {
	router := newRouter()
	router.Post("/auth", Auth("secret"))
	runAPITests(t, router, []apiTestCase{
		{"t1 - successful login", "POST", "/auth", `{"username":"demo", "password":"pass"}`, http.StatusOK, ""},
		{"t2 - unsuccessful login", "POST", "/auth", `{"username":"demo", "password":"bad"}`, http.StatusUnauthorized, ""},
	})
}