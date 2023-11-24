// Package routeHandling implements the RouteHandler for responding to requests
// to the route specified in each method's name
package routeHandling

import (
	"mathtestr.com/server/internal/dbHandler"
)

// RouteHandler responds to requests to the route specified in each method's name
type RouteHandler struct {
	dbHandler *dbHandler.DBHandler
}

// InitRouteHandler is constructor for RouteHandler, takes in pointer to initialized
// dbHandler and returns pointer to initializes RouteHandler
func InitRouteHandler(dbHandler *dbHandler.DBHandler) *RouteHandler {
	newRouteHandler := RouteHandler{
		dbHandler: dbHandler,
	}
	return &newRouteHandler
}
