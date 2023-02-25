package bootstrap

import (
	"gos/pkg/route"
	"gos/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	r := mux.NewRouter()
	routes.RegisterWebRoutes(r)

	route.SetRoute(r)

	return r
}
