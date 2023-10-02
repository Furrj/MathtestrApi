package types

type UserClientData struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

type SessionData struct {
	ID      int64  `json:"user_id"`
	UUID    string `json:"uuid"`
	Expires string `json:"expires"`
}
