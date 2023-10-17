package routeHandling

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
)

func (r *RouteHandler) Login(ctx *gin.Context) {
	var loginPayload schemas.LoginPayload
	loginResponse := schemas.LoginResponse{
		Valid: false,
	}

	if err := ctx.BindJSON(&loginPayload); err != nil {
		fmt.Printf("Error binding json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("%+v\n", loginPayload)

	valid, err := checkLoginInfo(r, loginPayload)
	if err != nil {
		fmt.Printf("Error in checkLoginInfo: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	if !valid {
		ctx.JSON(http.StatusOK, loginResponse)
		return
	}
	loginResponse.Valid = true
	ctx.JSON(http.StatusOK, loginResponse)
}

func checkLoginInfo(r *RouteHandler, loginPayload schemas.LoginPayload) (bool, error) {
	exists, err := r.dbHandler.CheckIfUsernameExists(loginPayload.Username)
	if err != nil {
		fmt.Printf("Error in CheckIfUsernameExists: %+v\n", err)
		return false, err
	}
	if !exists {
		return false, nil
	}

	userData, err := r.dbHandler.GetUserByUsername(loginPayload.Username)
	if err != nil {
		fmt.Printf("Error in GetUserByUsername: %+v\n", err)
		return false, err
	}

	if userData.Username == loginPayload.Username && userData.Password == loginPayload.Password {
		return true, nil
	}
	return false, nil
}
