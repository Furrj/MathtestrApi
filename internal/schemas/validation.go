package schemas

type ValidationPayload struct {
	ID         uint32 `json:"user_id"`
	SessionKey string `json:"session_key"`
}
type ValidationResponse struct {
	Valid       bool                  `json:"valid"`
	UserData    BasicUserDataResponse `json:"user_data,omitempty"`
	StudentData StudentDataResponse   `json:"student_data,omitempty"`
	TeacherData TeacherDataResponse   `json:"teacher_data,omitempty"`
}
