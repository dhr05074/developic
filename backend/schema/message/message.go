package message

type ProblemMessage struct {
	ProblemID string  `json:"problem_id"`
	SubmitID  *string `json:"submit_id"`
}
