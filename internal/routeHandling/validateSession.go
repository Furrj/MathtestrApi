package routeHandling

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
)

func (r *RouteHandler) ValidateSession(ctx *gin.Context) {
	var validationPayload schemas.ValidationPayload
	validationResponse := schemas.ValidationResponse{
		Valid: false,
	}

	// Bind request body
	if err := ctx.BindJSON(&validationPayload); err != nil {
		fmt.Printf("Error binding validationPayload json: %+v\n", err)
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
		fmt.Printf("Username '%s' could not be found", sessionData.Username)
		ctx.JSON(http.StatusOK, validationResponse)
		return
	}

	userData, err := r.dbHandler.GetBasicUserInfoByUsername(sessionData.Username)
	if err != nil {
		fmt.Printf("Error searching for user data: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
	}
	validationResponse.UserData = userData

	userSessionData, err := r.dbHandler.GetSessionDataByUserID(int(userData.ID))
	if err != nil {
		fmt.Printf("Error getting user session data: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
	}

	if userSessionData.SessionKey != validationPayload.SessionKey {
		ctx.JSON(http.StatusOK, validationResponse)
		return
	}

	// Get role-related data
	switch userData.Role {
	case "S":
		studentData, err := r.dbHandler.GetAllStudentDataByUsername(userData.Username)
		if err != nil {
			fmt.Printf("Error retrieving student data: %+v\n", err)
			ctx.JSON(http.StatusOK, validationResponse)
			return
		}
		validationResponse.StudentData.Period = studentData.Period
		validationResponse.StudentData.TeacherId = studentData.TeacherID
	case "T":
		teacherData, err := r.dbHandler.GetAllTeacherDataByUsername(userData.Username)
		if err != nil {
			fmt.Printf("Error retrieving teacher data: %+v\n", err)
			ctx.JSON(http.StatusOK, validationResponse)
			return
		}
		loginResponse.TeacherData.Periods = teacherData.Periods
	}

	validationResponse.Valid = true
	ctx.JSON(http.StatusOK, validationResponse)
}
