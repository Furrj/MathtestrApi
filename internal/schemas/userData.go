package schemas

type UserClientData struct {
	ID         uint32 `json:"user_id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	Period     uint8  `json:"period"`
	SessionKey string `json:"session_key"`
}

type SessionData struct {
	ID         uint32 `json:"user_id"`
	SessionKey string `json:"session_key"`
	Expires    uint64 `json:"expires"`
}

type AllUserData struct {
	ID         uint32 `json:"user_id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	Period     uint8  `json:"period"`
	Password   string `json:"password"`
	Firstname  string `json:"first_name"`
	Lastname   string `json:"last_name"`
	SessionKey string `json:"session_key"`
	Expires    uint64 `json:"expires"`
}
