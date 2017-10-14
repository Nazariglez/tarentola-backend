// Created by nazarigonzalez on 4/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"net/http"
)

func isLogged(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := controllers.GetRequestUserID(r)
		if id == 0 {
			controllers.SendBadRequest(w, r, controllers.GetRequestAuthError(r))
			return
		}

		next(w, r)
	}
}
