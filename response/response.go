package response

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
	Total   int         `json:"total,omitempty"`
}
