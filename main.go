package main

import (
	"net/http"

	_ "github.com/zhanbai/gos/router"
)

func main() {
	http.ListenAndServe(":8080", nil)
}
