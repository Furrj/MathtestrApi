package schemas

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Valid bool           `json:"valid"`
	User  UserClientData `json:"user"`
}
