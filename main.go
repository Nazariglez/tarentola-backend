// Created by nazarigonzalez on 30/9/17.

package main

import (
	"fmt"
	"net/http"

	"github.com/nazariglez/tarentola-backend/api/router"
	"github.com/nazariglez/tarentola-backend/database"
)

func main() {
	err := database.Open(false, false)
	if err != nil {
		panic(err)
	}

	defer database.Close()

	http.Handle("/", router.GetRouter())
	fmt.Println("Listening on 127.0.0.1:8000...")
	fmt.Errorf("%+v", http.ListenAndServe(":8000", nil))
}
