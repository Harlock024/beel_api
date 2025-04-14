package responses

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Status      bool   `json:"status"`
}

type TaskResponses []TaskResponse
