// Created by nazarigonzalez on 4/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"net/http"
)

func isLogged(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := controllers.ValidateToken(controllers.GetToken(r))
		if err != nil {
			controllers.SendForbidden(w, r, err.Error())
			return
		}

		next(w, r)
	}
}
