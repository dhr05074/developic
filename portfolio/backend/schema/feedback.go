package schema

type Feedback struct {
	Subject string `json:"subject"`
	Score   uint   `json:"score"`
	Comment string `json:"comment"`
}

type Recommendation struct {
	Subject string `json:"subject"`
	Score   uint   `json:"score"`
	Reason  string `json:"reason"`
}
