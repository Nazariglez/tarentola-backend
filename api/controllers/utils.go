// Created by nazarigonzalez on 14/10/17.

package controllers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/nazariglez/tarentola-backend/config"
	"net/http"
	"strings"
)

func GetToken(r *http.Request) string {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Bearer" {
		return ""
	}

	return auth[1]
}

func ValidateToken(token string) (*AuthClaims, error) {
	t, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(config.Data.Auth.Secret), nil
	})
	if err != nil {
		return &AuthClaims{}, err
	}

	if claims, ok := t.Claims.(*AuthClaims); ok && t.Valid {
		return claims, nil
	}

	return &AuthClaims{}, errors.New("Invalid token")
}

func GetRequestUserID(r *http.Request) uint {
	return r.Context().Value("userID").(uint)
}

func GetRequestID(r *http.Request) string {
	return r.Context().Value("rid").(string)
}

func GetRequestAuthError(r *http.Request) string {
	return r.Context().Value("authErr").(string)
}
