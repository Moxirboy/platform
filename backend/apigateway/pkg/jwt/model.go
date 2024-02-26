package jwt

import "github.com/gbrlsnchs/jwt/v3"

type TokenMetadata struct {
	jwt.Payload
	Role string `json:"role"`
}
