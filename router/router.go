package router

import (
	"net/http"

	"github.com/zhanbai/gos/api/controller/hello"
)

func init() {
	http.HandleFunc("/hello", hello.Get)
}
