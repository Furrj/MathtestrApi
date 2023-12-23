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
	fmt.Printf("%+v\n", validationPayload)

	userSessionData, err := r.dbHandler.GetSessionDataByUserID(int(validationPayload.ID))
	if err != nil {
		fmt.Printf("Error getting user session data: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
	}

	if userSessionData.SessionKey != validationPayload.SessionKey {
		ctx.JSON(http.StatusOK, validationResponse)
		return
	}

	// Get basic user info
	userData, err := r.dbHandler.GetBasicUserInfoByID(validationPayload.ID)
	if err != nil {
		fmt.Printf("Error searching for user data: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
	}
	validationResponse.UserData.ID = userData.ID
	validationResponse.UserData.Username = userData.Username
	validationResponse.UserData.FirstName = userData.FirstName
	validationResponse.UserData.LastName = userData.LastName
	validationResponse.UserData.Role = userData.Role

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
		validationResponse.TeacherData.Periods = teacherData.Periods
	}

	validationResponse.Valid = true
	ctx.JSON(http.StatusOK, validationResponse)
}
