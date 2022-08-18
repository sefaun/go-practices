package types

type LoginResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
