package userHandlers

import (
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"mathtestr.com/server/internal/dbHandling"
	"mathtestr.com/server/internal/schemas"
)

func CreateNewUser(db *pgx.Conn, registerPayload schemas.RegisterPayload) (schemas.UserClientData, error) {
	var userClientData schemas.UserClientData

	// Insert user using registerPayload
	if err := dbHandling.InsertUser(db, registerPayload); err != nil {
		return userClientData, err
	}

	// Get ID from created, user/doublecheck insert was correct
	createdUser, err := dbHandling.FindByUsername(db, registerPayload.Username)
	if err != nil {
		log.Print("Error in FindByUsername")
		return userClientData, err
	}

	// Create user session data
	userSessionData := schemas.SessionData{
		ID:      createdUser.ID,
		UUID:    uuid.New().String(),
		Expires: strconv.FormatInt(time.Now().Unix()+7776000, 10),
	}

	if err := dbHandling.InsertSessionData(db, userSessionData); err != nil {
		return userClientData, err
	}

	userClientData.ID = userSessionData.ID
	userClientData.UUID = userSessionData.UUID
	userClientData.Username = registerPayload.Username

	return userClientData, nil
}
