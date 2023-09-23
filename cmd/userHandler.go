package main

import (
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func CreateNewUser(db *pgx.Conn, registerPayload RegisterPayload) (UserClientData, error) {
	var userClientData UserClientData

	// Insert user using registerPayload
	if err := InsertUser(db, registerPayload); err != nil {
		return userClientData, err
	}

	// Get ID from created, user/doublecheck insert was correct
	createdUser, err := FindByUsername(db, registerPayload.Username)
	if err != nil {
		log.Print("Error in FindByUsername")
		return userClientData, err
	}

	// Create user session data
	userSessionData := SessionData{
		ID:      createdUser.ID,
		UUID:    uuid.New().String(),
		Expires: strconv.FormatInt(time.Now().Unix()+7776000, 10),
	}

	if err := InsertSessionData(db, userSessionData); err != nil {
		return userClientData, err
	}

	userClientData.ID = userSessionData.ID
	userClientData.UUID = userSessionData.UUID
	userClientData.Username = registerPayload.Username

	return userClientData, nil
}
