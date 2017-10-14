// Created by nazarigonzalez on 15/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"net/http"
)

func ParseForm(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			controllers.SendServerError(w, r, err)
			return
		}

		next.ServeHTTP(w, r)
	}
}
