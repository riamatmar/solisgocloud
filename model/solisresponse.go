package model

type SolisResponse struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Data    string `json:"data"`
}
