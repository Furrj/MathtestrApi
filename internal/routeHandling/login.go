package routeHandling

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
)

func (r *RouteHandler) Login(ctx *gin.Context) {
	var loginPayload schemas.LoginPayload

	if err := ctx.BindJSON(&loginPayload); err != nil {
		fmt.Printf("Error binding json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("%+v\n", loginPayload)
}
