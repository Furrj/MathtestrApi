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
	Valid bool           `json:"valid"`
	User  UserClientData `json:"user"`
}
