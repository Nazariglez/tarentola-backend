// Created by nazarigonzalez on 15/10/17.

package middlewares

import (
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"github.com/nazariglez/tarentola-backend/logger"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := controllers.GetRequestID(r)
		user := controllers.GetRequestUserID(r)
		logger.Log.Tracef("[User:%d - %s] Request (%s) - %s %s", user, r.RemoteAddr, rid, r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}
