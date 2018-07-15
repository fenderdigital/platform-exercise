package router

import (
	"testing"
)

func TestRoutes(t *testing.T) {
	var r = Routes()
	t.Log(r)
}

// func TestLoginHandler(t *testing.T) {
// 	var login = LoginHandler(w, r)
// }
