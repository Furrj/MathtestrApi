package userHandlers

import (
	"time"

	"github.com/google/uuid"
	"mathtestr.com/server/internal/schemas"
)

func GenerateNewUserSessionData(id int) schemas.SessionData {
	userSessionData := schemas.SessionData{
		ID:      uint32(id),
		UUID:    uuid.New().String(),
		Expires: uint64(time.Now().Unix() + 7776000),
	}
	return userSessionData
}
