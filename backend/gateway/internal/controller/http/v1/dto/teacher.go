package dto


type Teacher struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdateAt string `json:"update_at"`
	DeletedAt string `json:"deleted_at"`
	Verified bool `json:"verified"`
}