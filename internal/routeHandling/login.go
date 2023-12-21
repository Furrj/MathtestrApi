package routeHandling

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
)

func (r *RouteHandler) Login(ctx *gin.Context) {
	var loginPayload schemas.LoginPayload

	// Bind loginPayload
	if err := ctx.BindJSON(&loginPayload); err != nil {
		fmt.Printf("Error binding json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("Login Payload: %+v\n", loginPayload)

	// Check info and get loginResponse
	loginResponse, err := checkLoginInfo(r, loginPayload)
	if err != nil {
		fmt.Printf("Error in checkLoginInfo: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	ctx.JSON(http.StatusOK, loginResponse)
}

func checkLoginInfo(r *RouteHandler, loginPayload schemas.LoginPayload) (schemas.LoginResponse, error) {
	loginResponse := schemas.LoginResponse{
		Valid: false,
	}

	// Check if username exists
	exists, err := r.dbHandler.CheckIfUsernameExists(loginPayload.Username)
	if err != nil {
		fmt.Printf("Error in CheckIfUsernameExists: %+v\n", err)
		return loginResponse, err
	}
	if !exists {
		return loginResponse, nil
	}

	// Get user data
	userData, err := r.dbHandler.GetBasicUserInfoByUsername(loginPayload.Username)
	if err != nil {
		fmt.Printf("Error in GetBasicUserInfoByUsername: %+v\n", err)
		return loginResponse, err
	}

	userClientData, err := r.dbHandler.GetSessionDataByUserID(int(userData.ID))
	if err != nil {
		fmt.Printf("Error in GetSessionDataByUserID: %+v\n", err)
		return loginResponse, err
	}

	if userData.Username == loginPayload.Username && userData.Password == loginPayload.Password {
		loginResponse.Valid = true
		loginResponse.User.ID = userData.ID
		loginResponse.User.Username = userData.Username
		loginResponse.User.FirstName = userData.FirstName
		loginResponse.User.LastName = userData.LastName
		loginResponse.User.Role = userData.Role
		loginResponse.User.SessionKey = userClientData.SessionKey
	}

	return loginResponse, nil
}
