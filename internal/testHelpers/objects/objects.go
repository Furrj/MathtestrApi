// Packate objects includes objects that can be shared across
// different unit tests
package objects

import "mathtestr.com/server/internal/schemas"

var TestRegisterPayload = schemas.RegisterPayload{
	Username:  "a",
	Password:  "a",
	FirstName: "Jackson",
	LastName:  "Furr",
	Period:    0,
	Teacher:   2,
}

var TestSessionData = schemas.SessionData{
	ID:         1,
	SessionKey: "test_uuid",
	Expires:    1234,
}

var TestResultsData = schemas.TestResults{
	ID:            1,
	Score:         100,
	Min:           0,
	Max:           12,
	QuestionCount: 10,
	Operations:    "multiplication,addition",
}

var TestAllUserData = schemas.AllUserData{
	Username:   "a",
	Password:   "a",
	FirstName:  "Jackson",
	LastName:   "Furr",
	Period:     0,
	Teacher:    2,
	Role:       "S",
	ID:         1,
	SessionKey: "test_uuid",
	Expires:    1234,
}

var TestLoginPayload = schemas.LoginPayload{
	Username: "a",
	Password: "a",
}
