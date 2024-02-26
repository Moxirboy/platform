package jwt

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)
var mySigningKey = []byte("secret") // Change this to your secret key

func createToken() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["authorized"] = true
    claims["user"] = "John Doe"
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

    tokenString, err := token.SignedString(mySigningKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}