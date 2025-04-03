package request

type QuestionnaireUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}
