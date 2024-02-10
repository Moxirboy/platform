package postgres


const (
	CreateUser = `INSERT INTO users (id, username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, username, email, created_at, updated_at`
)