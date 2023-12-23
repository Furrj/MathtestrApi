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
