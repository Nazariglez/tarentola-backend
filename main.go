// Created by nazarigonzalez on 30/9/17.

package main

import (
	"net/http"

	"github.com/nazariglez/tarentola-backend/api/router"
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/content"
	"github.com/nazariglez/tarentola-backend/database"
	"github.com/nazariglez/tarentola-backend/email"
	"github.com/nazariglez/tarentola-backend/logger"
	"strconv"
)

func main() {
	if config.Data.IsProd() {
		logger.Log.Infof("%s initiated in 'production' mode.", config.Data.Name)
	} else {
		logger.Log.Debugf("%s initiated in '%s' mode.", config.Data.Name, config.Data.Environment)
	}

	err := database.Open()
	if err != nil {
		logger.Log.Fatal(err)
	}

	defer database.Close()

	go content.Serve()

	e := email.EmailBCC{
		To:   []string{"nazari.nz@gmail.com", "nazari.nz+test@gmail.com"},
		Body: email.ConfirmEmailTemplate,
		Data: map[string]interface{}{
			"name":              "Manuel",
			"confirmation_link": "http://google.es",
		},
		Subject: "Confirmation link...",
	}

	if err := e.Send(); err != nil {
		panic(err)
	}

	port := ":" + strconv.Itoa(config.Data.Port)
	handler := router.AllowCORS(router.GetRouter())

	logger.Log.Logf("Serving API on 127.0.0.1%s...", port)
	logger.Log.Fatal(http.ListenAndServe(port, handler))
}
