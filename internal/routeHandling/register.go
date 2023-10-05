package routeHandling

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/logger"
	"mathtestr.com/server/internal/schemas"
	"mathtestr.com/server/internal/userHandlers"
)

/*
POST("/register")
Recieves: RegisterPayload
Sends: RegisterResponse or ErrorCode
*/
func (r *RouteHandler) Register(ctx *gin.Context) {
	var registerPayload schemas.RegisterPayload

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
		ctx.String(http.StatusNotFound, "Username already exists")
		return
	}

	// Insert into user_info
	if err := r.dbHandler.InsertUser(registerPayload); err != nil {
		fmt.Printf("Error inserting user from /register: %+v\n", err)
		ctx.String(http.StatusNotFound, "Username already exists")
		return
	}

	// Get ID from newly inserted user
	id, err := r.dbHandler.GetUserIDByUsername(registerPayload.Username)
	if err != nil {
		fmt.Printf("Error searching for newly inserted user: %+v\n", err)
	}

	// Generate and insert Session Data
	sessionData := userHandlers.GenerateNewUserSessionData(id)
	if err := r.dbHandler.InsertSessionData(sessionData); err != nil {
		fmt.Printf("Error inserting session data: %+v\n", err)
	}

	// Generate and send response
	userClientData := schemas.UserClientData{
		ID:         sessionData.ID,
		Username:   registerPayload.Username,
		SessionKey: sessionData.SessionKey,
	}
	registerResponse := schemas.RegisterResponse{
		Valid: true,
		User:  userClientData,
	}

	ctx.JSON(http.StatusOK, registerResponse)

	// Log and backup
	go logger.WriteLog("POST /register", fmt.Sprintf("%+v\n", registerPayload))
	cmd := exec.Command("./backup.sh")
	if err := cmd.Run(); err != nil {
		log.Printf("Error backing up Postgres: %+v\n", err)
	}
}