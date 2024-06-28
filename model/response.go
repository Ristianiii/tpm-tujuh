package model

type ProductResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// type Response struct {
// 	Error   string      `json:"error,omitempty"`
// 	Success bool        `json:"success"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }
