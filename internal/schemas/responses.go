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
