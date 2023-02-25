package route

import (
	"github.com/gorilla/mux"
)

var route *mux.Router

func SetRoute(r *mux.Router) {
	route = r
}

// Name2URL 根据路由名称生成 URL
func Name2URL(name string, pairs ...string) string {
	url, err := route.Get(name).URL(pairs...)
	if err != nil {
		return ""
	}

	return url.String()
}
