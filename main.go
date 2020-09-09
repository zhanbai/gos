package main

import (
	"fmt"
	"net/http"

	"github.com/zhanbai/gos/library/config"
	_ "github.com/zhanbai/gos/router"
)

func main() {
	config := config.Get()
	address := config.Server.Address
	fmt.Printf("Listening and serving HTTP on %s\n", address)
	http.ListenAndServe(address, nil)
}
