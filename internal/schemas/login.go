package schemas

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Valid       bool                  `json:"valid"`
	UserData    BasicUserDataResponse `json:"user_data,omitempty"`
	StudentData StudentDataResponse   `json:"student_data,omitempty"`
	TeacherData TeacherDataResponse   `json:"teacher_data,omitempty"`
	SessionKey  string                `json:"session_key"`
}
