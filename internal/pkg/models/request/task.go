package request

type TaskRequest struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
