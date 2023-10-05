package routeHandling

import (
	"mathtestr.com/server/internal/dbHandling"
)

type RouteHandler struct {
	dbHandler *dbHandling.DBHandler
}

func InitRouteHandler(dbHandler *dbHandling.DBHandler) *RouteHandler {
	newRouteHandler := RouteHandler{
		dbHandler: dbHandler,
	}
	return &newRouteHandler
}
