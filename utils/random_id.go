// Created by nazarigonzalez on 18/10/17.

package utils

import "math/rand"

const alphaNumCharacters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ"

func GetRandomID(n int) string {
	id := make([]byte, n)
	for i, _ := range id {
		id[i] = alphaNumCharacters[rand.Intn(len(alphaNumCharacters))]
	}

	return string(id)
}
