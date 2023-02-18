package main

import (
	"gos/bootstrap"
	"net/http"
)

func main() {
	r := bootstrap.SetupRoute()
	http.ListenAndServe(":3000", r)
}
