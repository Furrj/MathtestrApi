package schemas

type RegisterPayload struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Period    uint8  `json:"period,string"`
	TeacherID uint32 `json:"teacher_id,string"`
}

type RegisterResponse struct {
	Valid       bool                  `json:"valid"`
	UserData    BasicUserDataResponse `json:"user_data,omitempty"`
	StudentData StudentDataResponse   `json:"student_data,omitempty"`
	TeacherData TeacherDataResponse   `json:"teacher_data,omitempty"`
	SessionKey  string                `json:"session_key"`
}
