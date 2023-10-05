package routeHandling

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
)

func (r *RouteHandler) ValidateSession(ctx *gin.Context) {
	var sessionData schemas.UserClientData

	// Bind request body
	if err := ctx.BindJSON(&sessionData); err != nil {
		fmt.Printf("Error binding json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("%+v\n", sessionData)
}
