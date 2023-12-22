package schemas

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Valid       bool          `json:"valid"`
	UserData    BasicUserData `json:"user_data,omitempty"`
	StudentData StudentData   `json:"student_data,omitempty"`
	TeacherData TeacherData   `json:"teacher_data,omitempty"`
	SessionKey  string        `json:"session_key"`
}
