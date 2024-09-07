package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/* //Generate a fake secret key
func init() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	stringB64 := base64.StdEncoding.EncodeToString(key)
	fmt.Printf("Secret key: %s", stringB64)

} */

func main() {
	config.Init()
	fmt.Printf("Secret key: %s\n", config.SecretKey)
	fmt.Printf("Server running on port %d", config.Port)

	r := router.Generate()
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port),
		r,
	))
}
