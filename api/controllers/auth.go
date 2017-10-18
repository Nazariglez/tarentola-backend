// Created by nazarigonzalez on 3/10/17.

package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/database/usermodel"
	"github.com/nazariglez/tarentola-backend/utils"
	"net/http"
	"strings"
	"time"
)

type AuthClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type loginObj struct {
	ID       uint   `json:"id"`
	ExpireAt int64  `json:"expireAt"`
	Token    string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	expireToken := time.Now().
		Add(time.Second * time.Duration(config.Data.Auth.TokenExpire)).
		Unix()

	email := strings.TrimSpace(r.Form.Get("email"))
	pass := strings.TrimSpace(r.Form.Get("password"))

	if email == "" || pass == "" {
		SendBadRequest(w, r, "Empty email or password.")
		return
	}

	if err := utils.ValidateEmailFormat(email); err != nil {
		SendBadRequest(w, r, err.Error())
		return
	}

	user, err := usermodel.FindToLogin(email, pass)
	if err != nil {
		SendServerError(w, r, err)
		return
	}

	if user == nil {
		SendBadRequest(w, r, "Invalid email or password.")
		return
	}

	if !user.IsActive() {
		if !user.Active {
			SendBadRequest(w, r, "This account must be activated first.")
			return
		}

		if user.Banned {
			msg := "Account banned."
			if user.BanTime != (time.Time{}) {
				msg = fmt.Sprintf("Account banned until: '%s'", user.BanTime.Format("02/01/2006"))
			}

			SendForbidden(w, r, msg)
			return
		}
	}

	claims := AuthClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    fmt.Sprintf("localhost:%d", config.Data.Port),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.Data.Auth.Secret))

	data := loginObj{
		ID:       user.ID,
		ExpireAt: expireToken,
		Token:    signedToken,
	}

	SendOk(w, r, data)
	return
}
