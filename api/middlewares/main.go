// Created by nazarigonzalez on 3/10/17.

package middlewares

import (
	"context"
	"github.com/nazariglez/tarentola-backend/api/controllers"
	"github.com/nazariglez/tarentola-backend/logger"
	"math/rand"
	"net/http"
)

const characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ"

type Middleware func(next http.HandlerFunc) http.HandlerFunc

var list = map[string][]Middleware{
	"isLogged": {
		isLogged,
	},
	"isNotLogged": {
		isNotLogged,
	},
}

func Apply(name string, controller http.HandlerFunc) http.HandlerFunc {
	mws, ok := list[name]
	if !ok {
		logger.Log.Fatalf("Invalid '%s' middleware name.", name)
		return nil
	}

	for i := len(mws) - 1; i >= 0; i-- {
		controller = mws[i](controller)
	}

	return controller
}

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
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func ParseForm(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			controllers.SendServerError(w, r, err)
			return
		}

		next(w, r)
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := controllers.GetRequestID(r)
		user := controllers.GetRequestUserID(r)
		logger.Log.Tracef("[User:%d - %s] Request (%s) - %s %s", user, r.RemoteAddr, rid, r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}
