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

type ProtoQuery struct {
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
	Page int 	`json:"page"`
	Size int 	`json:"size"`
	HasMore bool 	`json:"has_more"`
	Users []*Teacher
}