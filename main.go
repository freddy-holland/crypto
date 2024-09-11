package main

import (
	"fmt"

	"fholl.net/auth/crypto"
)

func main() {
	header := map[string]interface{}{
		"alg": "HS256", "typ": "JWT",
	}
	payload := map[string]interface{}{
		"exp": 1726134237, "sub": "123456789",
	}

	jwt := crypto.CreateJWT(header, payload)
	fmt.Println(jwt)
}
