package schemas

type TestResults struct {
	ID            uint32 `json:"user_id"`
	Score         uint8  `json:"score"`
	Min           int32  `json:"min"`
	Max           int32  `json:"max"`
	QuestionCount uint32 `json:"question_count"`
	Operations    string `json:"operations"`
	Timestamp     uint64 `json:"timestamp"`
}
