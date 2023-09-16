package main

type User struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
}

type RegisterPayload struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type RegisterResponse struct {
	Valid bool `json:"valid"`
	User  User `json:"user"`
}
