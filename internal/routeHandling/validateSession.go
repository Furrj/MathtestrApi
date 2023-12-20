package routeHandling

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
)

func (r *RouteHandler) ValidateSession(ctx *gin.Context) {
	var sessionData schemas.UserClientData
	validationResponse := schemas.SessionValidationResponse{
		Valid: false,
	}

	// Bind request body
	if err := ctx.BindJSON(&sessionData); err != nil {
		fmt.Printf("Error binding json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("%+v\n", sessionData)

	exists, err := r.dbHandler.CheckIfUsernameExists(sessionData.Username)
	if err != nil {
		fmt.Printf("Error searching for name: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	if !exists {
		fmt.Printf("%q could not be found in database", sessionData.Username)
		ctx.JSON(http.StatusOK, validationResponse)
		return
	}

	userData, err := r.dbHandler.GetBasicUserInfoByUsername(sessionData.Username)
	if err != nil {
		fmt.Printf("Error searching for user data: %+v\n", err)
		return
	}
	if userData.SessionKey != sessionData.SessionKey {
		ctx.JSON(http.StatusOK, validationResponse)
		return
	}
	validationResponse.Valid = true
	ctx.JSON(http.StatusOK, validationResponse)
}
