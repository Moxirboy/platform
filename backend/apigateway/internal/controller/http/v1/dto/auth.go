package dto

type LoginRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`	
	Password string `json:"password"`
}
type Response struct {
	Token string `json:"token"`
	Role string `json:"role"`
}

type SignUpRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Class     string `json:"class"`
}
type SignUpResponse struct {
	Id      string `json:"id"`
	Role    string `json:"role"`
	ClassId string `json:"class_id"`
}
