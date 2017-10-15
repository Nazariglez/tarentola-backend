// Created by nazarigonzalez on 15/10/17.

package middlewares

import (
	"context"
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"github.com/nazariglez/tarentola-backend/utils"
	"math/rand"
	"net/http"
)

const characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ"

func getRandomID() string {
	id := make([]byte, 10)
	for i, _ := range id {
		id[i] = characters[rand.Intn(len(characters))]
	}

	return string(id)
}

func InitRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "rid", getRandomID())

		claims, err := controllers.ValidateToken(controllers.GetToken(r))
		if err != nil {
			ctx = context.WithValue(ctx, "authErr", err.Error())
		}

		ctx = context.WithValue(ctx, "userID", claims.ID)
		ctx = context.WithValue(ctx, "ipAddr", utils.GetIPAddr(r))
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
