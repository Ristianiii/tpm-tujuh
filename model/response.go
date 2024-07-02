package model

type ProductResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type UserResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
