// Packate objects includes objects that can be shared across
// different unit tests
package objects

import "mathtestr.com/server/internal/schemas"

var TestRegisterPayloadStudent = schemas.RegisterPayload{
	Username:  "a",
	Password:  "a",
	FirstName: "Jackson",
	LastName:  "Furr",
	Period:    2,
	TeacherID: 1,
}

var TestSessionDataStudent = schemas.SessionData{
	ID:         2,
	SessionKey: "test_uuid",
	Expires:    1234,
}

var TestResultsDataStudent = schemas.TestResults{
	ID:            2,
	Score:         100,
	Min:           0,
	Max:           12,
	QuestionCount: 10,
	Operations:    "multiplication,addition",
}

var TestAllUserDataStudent = schemas.AllUserData{
	Username:   "a",
	Password:   "a",
	FirstName:  "Jackson",
	LastName:   "Furr",
	Period:     2,
	TeacherID:  1,
	Role:       "S",
	ID:         2,
	SessionKey: "test_uuid",
	Expires:    1234,
}

var TestLoginPayloadStudent = schemas.LoginPayload{
	Username: "a",
	Password: "a",
}

var TestRegisterPayloadTeacher = schemas.RegisterPayload{
	Username:  "MFurr",
	Password:  "password",
	FirstName: "Michelle",
	LastName:  "Furr",
	Period:    0,
	TeacherID: 1,
}

var TestTeacherInfo = schemas.TeacherData{
	ID:      1,
	Periods: 8,
}