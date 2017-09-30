// Created by nazarigonzalez on 30/9/17.

package main

import (
  "github.com/nazariglez/tarentola-backend/database"
)

func main() {
  err := database.Open(false, false)
  if err != nil {
    panic(err)
  }

  defer database.Close()
}