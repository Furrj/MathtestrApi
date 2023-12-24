package schemas

type BasicUserDataResponse struct {
	ID        uint32 `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type StudentDataResponse struct {
	TeacherId uint32 `json:"teacher_id"`
	Period    uint8  `json:"period"`
}

type TeacherDataResponse struct {
	Periods uint8 `json:"periods"`
}
type TestResultsResponse struct {
	Score         uint8  `json:"score"`
	Min           int32  `json:"min"`
	Max           int32  `json:"max"`
	QuestionCount uint32 `json:"question_count"`
	Operations    string `json:"operations"`
	Timestamp     uint64 `json:"timestamp"`
}

type ProfilePageResponse struct {
	Tests       []TestResultsResponse `json:"tests"`
	TeacherName string                `json:"teacher_name"`
}

type LoginResponse struct {
	Valid       bool                  `json:"valid"`
	UserData    BasicUserDataResponse `json:"user_data,omitempty"`
	StudentData StudentDataResponse   `json:"student_data,omitempty"`
	TeacherData TeacherDataResponse   `json:"teacher_data,omitempty"`
	SessionKey  string                `json:"session_key"`
}

type RegisterResponse struct {
	Valid       bool                  `json:"valid"`
	UserData    BasicUserDataResponse `json:"user_data,omitempty"`
	StudentData StudentDataResponse   `json:"student_data,omitempty"`
	TeacherData TeacherDataResponse   `json:"teacher_data,omitempty"`
	SessionKey  string                `json:"session_key"`
}

type ValidationResponse struct {
	Valid       bool                  `json:"valid"`
	UserData    BasicUserDataResponse `json:"user_data,omitempty"`
	StudentData StudentDataResponse   `json:"student_data,omitempty"`
	TeacherData TeacherDataResponse   `json:"teacher_data,omitempty"`
}
