// Created by nazarigonzalez on 15/10/17.

package policies

import (
	"github.com/nazariglez/tarentola-backend/api/middlewares"
	"github.com/nazariglez/tarentola-backend/logger"
	"net/http"
)

var policiesList = map[string][]middlewares.Middleware{
	"isLogged": {
		isLogged,
	},
	"isNotLogged": {
		isNotLogged,
	},
}

func Apply(name string, controller http.HandlerFunc) http.HandlerFunc {
	mws, ok := policiesList[name]
	if !ok {
		logger.Log.Fatalf("Invalid '%s' middleware name.", name)
		return nil
	}

	for i := len(mws) - 1; i >= 0; i-- {
		controller = mws[i](controller)
	}

	return controller
}
