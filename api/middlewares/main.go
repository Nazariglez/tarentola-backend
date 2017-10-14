// Created by nazarigonzalez on 3/10/17.

package middlewares

import (
	"net/http"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc
