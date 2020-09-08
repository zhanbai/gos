package router

import (
	"net/http"

	"github.com/zhanbai/go-skeleton/api/controller/hello"
)

func init() {
	http.HandleFunc("/hello", hello.Get)
}
