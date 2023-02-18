package bootstrap

import (
	"gos/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	r := mux.NewRouter()
	routes.RegisterWebRoutes(r)
	return r
}
