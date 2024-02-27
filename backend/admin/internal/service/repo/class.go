package repo

import (
	"context"
)

type ClassRepository interface {
	GetClassExist(ctx context.Context, className string) (bool, error)
	GetClassPassword(ctx context.Context, className string) (string, string, error)
	LinkClassUser(ctx context.Context, classID string, userID string) error
}
