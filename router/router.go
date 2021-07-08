package router

import (
	"net/http"

	"gos/api/controller/hello"
)

func init() {
	http.HandleFunc("/hello", hello.Get)
}
