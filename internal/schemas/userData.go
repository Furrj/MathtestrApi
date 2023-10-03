package schemas

type UserClientData struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	UUID     uint64 `json:"uuid"`
}

type SessionData struct {
	ID      int64  `json:"user_id"`
	UUID    string `json:"uuid"`
	Expires uint64 `json:"expires"`
}

type AllUserData struct {
	ID        int64  `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	UUID      string `json:"uuid"`
	Expires   uint64 `json:"expires"`
}
