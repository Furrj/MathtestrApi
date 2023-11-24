package routeHandling

import (
	"mathtestr.com/server/internal/dbHandler"
)

type RouteHandler struct {
	dbHandler *dbHandler.DBHandler
}

func InitRouteHandler(dbHandler *dbHandler.DBHandler) *RouteHandler {
	newRouteHandler := RouteHandler{
		dbHandler: dbHandler,
	}
	return &newRouteHandler
}
