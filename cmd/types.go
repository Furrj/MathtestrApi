package main

type UserClientData struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

type RegisterPayload struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type RegisterResponse struct {
	Valid bool           `json:"valid"`
	User  UserClientData `json:"user"`
}

type SessionData struct {
	ID      int64  `json:"user_id"`
	UUID    string `json:"uuid"`
	Expires string `json:"expires"`
}
