package routeHandling

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
)

func (r *RouteHandler) Login(ctx *gin.Context) {
	var loginPayload schemas.LoginPayload
	var loginResponse = schemas.LoginResponse{
		Valid: false,
	}

	// Bind loginPayload
	if err := ctx.BindJSON(&loginPayload); err != nil {
		fmt.Printf("Error binding json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("Login Payload: %+v\n", loginPayload)

	// Check if username exists
	exists, err := r.dbHandler.CheckIfUsernameExists(loginPayload.Username)
	if err != nil {
		fmt.Printf("Error in CheckIfUsernameExists: %+v\n", err)
		ctx.JSON(http.StatusNotFound, loginResponse)
		return
	}
	if !exists {
		fmt.Printf("Username '%s' does not exist", loginPayload.Username)
		ctx.JSON(http.StatusOK, loginResponse)
		return
	}

	// Get basic user data
	userData, err := r.dbHandler.GetBasicUserInfoByUsername(loginPayload.Username)
	if err != nil {
		fmt.Printf("Error in GetBasicUserInfoByUsername: %+v\n", err)
		ctx.JSON(http.StatusNotFound, loginResponse)
		return
	}
	loginResponse.UserData.ID = userData.ID
	loginResponse.UserData.Username = userData.Username
	loginResponse.UserData.FirstName = userData.FirstName
	loginResponse.UserData.LastName = userData.LastName
	loginResponse.UserData.Role = userData.Role

	// Get session data
	sessionData, err := r.dbHandler.GetSessionDataByUserID(int(userData.ID))
	if err != nil {
		fmt.Printf("Error retrieving session info: %+v\n", err)
		ctx.JSON(http.StatusNotFound, loginResponse)
		return
	}
	loginResponse.SessionKey = sessionData.SessionKey

	// Get role-related data
	switch userData.Role {
	case "S":
		studentData, err := r.dbHandler.GetAllStudentDataByUsername(loginPayload.Username)
		if err != nil {
			fmt.Printf("Error retrieving student data: %+v\n", err)
			ctx.JSON(http.StatusOK, loginResponse)
			return
		}
		loginResponse.StudentData.Period = studentData.Period
		loginResponse.StudentData.TeacherId = studentData.TeacherID
		loginResponse.Valid = true
	case "T":
		teacherData, err := r.dbHandler.GetAllTeacherDataByUsername(loginPayload.Username)
		if err != nil {
			fmt.Printf("Error retrieving teacher data: %+v\n", err)
			ctx.JSON(http.StatusOK, loginResponse)
			return
		}
		loginResponse.TeacherData.Periods = teacherData.Periods
		loginResponse.Valid = true
	}

	ctx.JSON(http.StatusOK, loginResponse)
}
