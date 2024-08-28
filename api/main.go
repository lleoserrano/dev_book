package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Init()
	fmt.Printf("Server running on port %d", config.Port)

	r := router.Generate()
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port),
		r,
	))
}
