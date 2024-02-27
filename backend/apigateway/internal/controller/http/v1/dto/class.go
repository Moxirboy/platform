package dto

type ClassRequest struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
}
type Class struct{
	Name string `json:"name"`
}