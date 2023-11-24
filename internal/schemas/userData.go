package schemas

type UserClientData struct {
	ID         uint32 `json:"user_id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	Period     uint8  `json:"period"`
	Teacher    uint   `json:"teacher"`
	SessionKey string `json:"session_key"`
}

type SessionData struct {
	ID         uint32 `json:"user_id"`
	SessionKey string `json:"session_key"`
	Expires    uint64 `json:"expires"`
}

type TestResults struct {
	ID            uint32 `json:"user_id"`
	Score         uint8  `json:"score"`
	Min           int32  `json:"min"`
	Max           int32  `json:"max"`
	QuestionCount uint32 `json:"question_count"`
	Operations    string `json:"operations"`
}

type AllUserData struct {
	ID         uint32 `json:"user_id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	Period     uint8  `json:"period"`
	Teacher    uint   `json:"teacher"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	SessionKey string `json:"session_key"`
	Expires    uint64 `json:"expires"`
}
