package main

import (
	"net/http"

	_ "github.com/zhanbai/go-skeleton/router"
)

func main() {
	http.ListenAndServe(":8080", nil)
}
