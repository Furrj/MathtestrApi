package schemas

type UserClientData struct {
	ID         uint32 `json:"user_id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Role       string `json:"role"`
	Period     uint8  `json:"period"`
	TeacherID  uint32 `json:"teacher_id"`
	SessionKey string `json:"session_key"`
}

type SessionData struct {
	ID         uint32 `json:"user_id"`
	SessionKey string `json:"session_key"`
	Expires    uint64 `json:"expires"`
}

type BasicUserData struct {
	ID        uint32 `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AllUserDataStudent struct {
	ID         uint32 `json:"user_id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	Period     uint8  `json:"period"`
	TeacherID  uint32 `json:"teacher_id"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	SessionKey string `json:"session_key"`
	Expires    uint64 `json:"expires"`
}

type AllUserDataTeacher struct {
	ID         uint32 `json:"user_id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	Periods    uint8  `json:"periods"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	SessionKey string `json:"session_key"`
	Expires    uint64 `json:"expires"`
}

type TeacherData struct {
	ID      uint32 `json:"user_id"`
	Periods uint8  `json:"periods"`
}

type StudentData struct {
	TeacherId uint32 `json:"teacher_id"`
	Period    uint8  `json:"period"`
}
