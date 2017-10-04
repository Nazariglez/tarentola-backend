// Created by nazarigonzalez on 4/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"net/http"
	"strings"
)

func isLogged(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Bearer" {
			controllers.SendForbidden(w, "Invalid auth token.")
			return
		}

		_, err := controllers.ValidateToken(auth[1])
		if err != nil {
			controllers.SendForbidden(w, err.Error())
			return
		}

		next(w, r)
	}
}
