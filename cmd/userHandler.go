package main

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func CreateNewUser(db *pgx.Conn, registerPayload RegisterPayload) (UserClientData, error) {
	var userClientData UserClientData

	InsertUser(db, registerPayload)
	user, err := FindByUsername(db, registerPayload.Username)
	if err != nil {
		return userClientData, err
	}

	userSessionData := SessionData{
		ID:      user.ID,
		UUID:    uuid.New().String(),
		Expires: strconv.FormatInt(time.Now().Unix()+7776000, 10),
	}

	userClientData.ID = user.ID
	userClientData.Username = user.Username
	userClientData.UUID = userSessionData.UUID

	return userClientData, nil
}
