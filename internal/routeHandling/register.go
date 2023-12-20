package routeHandling

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/schemas"
	"mathtestr.com/server/internal/userHandlers"
	"net/http"
)

// Register recieves a RegisterPayload, then checks if the username is valid,
// inserts into user_info, generates SessionData, then sends a RegisterResponse
func (r *RouteHandler) Register(ctx *gin.Context) {
	var registerPayload schemas.RegisterPayload
	var registerResponse schemas.RegisterResponse
	registerResponse.Valid = false

	// Bind request body
	if err := ctx.BindJSON(&registerPayload); err != nil {
		fmt.Printf("Error binding json: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	fmt.Printf("%+v\n", registerPayload)

	// Check if username currently exists
	exists, err := r.dbHandler.CheckIfUsernameExists(registerPayload.Username)
	if err != nil {
		fmt.Printf("Error checking username validity: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error")
		return
	}
	if exists {
		ctx.JSON(http.StatusOK, registerResponse)
		return
	}

	// Insert into user_info
	if err := r.dbHandler.InsertUserInfo("S", registerPayload); err != nil {
		fmt.Printf("Error inserting user from /register: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error inserting new user")
		return
	}

	// Get ID from newly inserted user
	id, err := r.dbHandler.GetUserIDByUsername(registerPayload.Username)
	if err != nil {
		fmt.Printf("Error searching for newly inserted user: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error inserting new user")
		return
	}

	// Insert role-related data
	// FIXME: Switch-case for possible roles
	if err := r.dbHandler.InsertStudentInfo(uint32(id), registerPayload); err != nil {
		fmt.Printf("Error inserting student info: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error inserting student info")
		return
	}

	// Generate and insert Session Data
	sessionData := userHandlers.GenerateNewUserSessionData(id)
	if err := r.dbHandler.InsertSessionData(sessionData); err != nil {
		fmt.Printf("Error inserting session data: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error generating and inserting session data")
		return
	}

	//Get all new user data
	userData, err := r.dbHandler.GetBasicUserInfoByUsername(registerPayload.Username)
	if err != nil {
		fmt.Printf("Error retrieving new user information: %+v\n", err)
		ctx.String(http.StatusNotFound, "Error retrieving new user data afer insertion")
		return
	}

	// Generate and send response
	userClientData := schemas.UserClientData{
		ID:         userData.ID,
		Username:   userData.Username,
		Role:       userData.Role,
		Period:     userData.Period,
		TeacherID:  userData.TeacherID,
		SessionKey: userData.SessionKey,
	}
	registerResponse.Valid = true
	registerResponse.User = userClientData

	ctx.JSON(http.StatusOK, registerResponse)

	// Backup
	//cmd := exec.Command("./backup.sh")
	//if err := cmd.Run(); err != nil {
	//	log.Printf("cError backing up Postgres: %+v\n", err)
	//}
}
