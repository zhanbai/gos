package main

import (
	"fmt"
	"net/http"

	"gos/library/config"
	_ "gos/router"
)

func main() {
	config := config.Get()
	address := config.Server.Address
	fmt.Printf("Listening and serving HTTP on %s\n", address)
	http.ListenAndServe(address, nil)
}
