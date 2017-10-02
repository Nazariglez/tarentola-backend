// Created by nazarigonzalez on 30/9/17.

package main

import (
	"net/http"

	"github.com/nazariglez/tarentola-backend/api/router"
	"github.com/nazariglez/tarentola-backend/config"
	"github.com/nazariglez/tarentola-backend/database"
	"github.com/nazariglez/tarentola-backend/logger"
	"strconv"
)

func main() {
	err := database.Open()
	if err != nil {
		logger.Log.Fatal(err)
	}

	defer database.Close()

	http.Handle("/", router.GetRouter())
	port := ":" + strconv.Itoa(config.Data.Port)
	logger.Log.Logf("Listening on 127.0.0.1%s...", port)
	logger.Log.Error(http.ListenAndServe(port, nil))
}
