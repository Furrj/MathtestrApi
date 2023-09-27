package userHandlers

import (
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"mathtestr.com/server/internal/dbHandlers"
	"mathtestr.com/server/internal/types"
)

func CreateNewUser(db *pgx.Conn, registerPayload types.RegisterPayload) (types.UserClientData, error) {
	var userClientData types.UserClientData

	// Insert user using registerPayload
	if err := dbHandlers.InsertUser(db, registerPayload); err != nil {
		return userClientData, err
	}

	// Get ID from created, user/doublecheck insert was correct
	createdUser, err := dbHandlers.FindByUsername(db, registerPayload.Username)
	if err != nil {
		log.Print("Error in FindByUsername")
		return userClientData, err
	}

	// Create user session data
	userSessionData := types.SessionData{
		ID:      createdUser.ID,
		UUID:    uuid.New().String(),
		Expires: strconv.FormatInt(time.Now().Unix()+7776000, 10),
	}

	if err := dbHandlers.InsertSessionData(db, userSessionData); err != nil {
		return userClientData, err
	}

	userClientData.ID = userSessionData.ID
	userClientData.UUID = userSessionData.UUID
	userClientData.Username = registerPayload.Username

	return userClientData, nil
}
