package schemas

type UserClientData struct {
	ID       uint32 `json:"user_id"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

type SessionData struct {
	ID      uint32 `json:"user_id"`
	UUID    string `json:"uuid"`
	Expires uint64 `json:"expires"`
}

type AllUserData struct {
	ID        uint32 `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	UUID      string `json:"uuid"`
	Expires   uint64 `json:"expires"`
}
