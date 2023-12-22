package schemas

type ValidationPayload struct {
	ID         uint32 `json:"user_id"`
	SessionKey string `json:"session_key"`
}
type ValidationResponse struct {
	Valid       bool          `json:"valid"`
	UserData    BasicUserData `json:"user_data,omitempty"`
	StudentData StudentData   `json:"student_data,omitempty"`
	TeacherData TeacherData   `json:"teacher_data,omitempty"`
}
