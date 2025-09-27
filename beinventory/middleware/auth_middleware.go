package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/util"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	helper.GoDoEnv()
	var Token = os.Getenv("Token")
	url := strings.TrimPrefix(r.URL.Path, "/api/")
	tokenn, err := r.Cookie(Token)
	if r.Method == "POST" {
		switch url {
		case "login", "pengguna", "logout":
			middleware.Handler.ServeHTTP(w, r)
			return
		default:
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "Mohon Untuk Login",
				}
				helper.WriteToResponse(w, webResponse)
				return
			}
			middleware.Handler.ServeHTTP(w, r)
			return
		}
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Cookie Tidak Ditemukan",
		}
		helper.WriteToResponse(w, webResponse)
		return
	}
	_, err = util.Decodetoken(tokenn.Value)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		helper.WriteToResponse(w, map[string]interface{}{
			"Code":    http.StatusBadRequest,
			"Status":  "Bad Request",
			"Message": err.Error(),
		})
		return
	}
	middleware.Handler.ServeHTTP(w, r)
}
