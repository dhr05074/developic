package schema

type Report struct {
	ProjectFeedbacks         []Feedback       `json:"project_feedbacks"`
	TechStackFeedbacks       []Feedback       `json:"tech_stack_feedbacks"`
	ProjectRecommendations   []Recommendation `json:"project_recommendations"`
	TechStackRecommendations []Recommendation `json:"tech_stack_recommendations"`
}
