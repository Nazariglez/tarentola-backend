// Created by nazarigonzalez on 4/10/17.

package policies

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"net/http"
	"strings"
)

func isNotLogged(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.TrimSpace(r.Header.Get("Authorization"))

		if auth != "" {
			controllers.SendForbidden(w, r, "You mustn't be logged.")
			return
		}

		next(w, r)
	}
}
