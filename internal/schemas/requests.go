package schemas

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Period    uint8  `json:"period,string"`
	TeacherID uint32 `json:"teacher_id,string"`
}

type ValidationPayload struct {
	ID         uint32 `json:"user_id"`
	SessionKey string `json:"session_key"`
}

type ProfilePagePayload struct {
	ID        uint32 `json:"user_id"`
	TeacherID uint32 `json:"teacher_id,string"`
}
