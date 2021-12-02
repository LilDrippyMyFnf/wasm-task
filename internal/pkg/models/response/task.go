package response

type TaskResponse struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TasksResponse struct {
	Tasks []*TaskResponse `json:"tasks"`
}
