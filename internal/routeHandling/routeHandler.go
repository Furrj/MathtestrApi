package routeHandling

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mathtestr.com/server/internal/dbHandling"
	"mathtestr.com/server/internal/schemas"
	"mathtestr.com/server/internal/userHandlers"
)

type RouteHandler struct {
	dbHandler *dbHandling.DBHandler
}

func InitRouteHandler(dbHandler *dbHandling.DBHandler) *RouteHandler {
	newRouteHandler := RouteHandler{
		dbHandler: dbHandler,
	}
	return &newRouteHandler
}

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
		ID:       sessionData.ID,
		Username: registerPayload.Username,
		UUID:     sessionData.UUID,
	}
	registerResponse := schemas.RegisterResponse{
		Valid: true,
		User:  userClientData,
	}
	ctx.JSON(http.StatusOK, registerResponse)
}

// func RegisterPost(db *pgx.Conn) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var registerPayload schemas.RegisterPayload
// 		var registerResponse schemas.RegisterResponse
// 		registerResponse.Valid = false

// 		// Marshall JSON from request body
// 		if err := ctx.BindJSON(&registerPayload); err != nil {
// 			log.Printf("Error binding register payload:\n%+v\n", err)
// 			ctx.String(http.StatusNotFound, "Error")
// 			return
// 		}
// 		fmt.Printf("%+v\n", registerPayload)

// 		// Check if username exists
// 		user, err := dbHandling.FindByUsername(db, registerPayload.Username)
// 		if err != nil {
// 			log.Print("Error in FindByUsername")
// 			ctx.String(http.StatusBadRequest, "Error in FindByUsername")
// 			return
// 		}

// 		// If username doesn't exist
// 		if user.ID == -1 {
// 			createdUserClientData, err := userHandlers.CreateNewUser(db, registerPayload)
// 			if err != nil {
// 				log.Print("Error in CreateUser")
// 				ctx.String(http.StatusBadRequest, "Error in CreateUser")
// 				return
// 			}

// 			registerResponse.Valid = true
// 			registerResponse.User = createdUserClientData
// 			ctx.JSON(http.StatusOK, registerResponse)
// 			return
// 		}
// 		ctx.String(http.StatusBadRequest, "Error")
// 	}
// }
