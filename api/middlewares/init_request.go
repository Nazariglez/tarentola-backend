// Created by nazarigonzalez on 15/10/17.

package middlewares

import (
	"context"
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"github.com/nazariglez/tarentola-backend/utils"
	"net/http"
)

func InitRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "rid", utils.GetRandomID(10))

		claims, err := controllers.ValidateToken(controllers.GetToken(r))
		if err != nil {
			ctx = context.WithValue(ctx, "authErr", err.Error())
		}

		ctx = context.WithValue(ctx, "userID", claims.ID)
		ctx = context.WithValue(ctx, "ipAddr", utils.GetIPAddr(r))
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
