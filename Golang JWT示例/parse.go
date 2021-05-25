/*
 * @Author: your name
 * @Date: 2021-04-08 15:04:32
 * @LastEditTime: 2021-04-08 15:46:48
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/Golang JWT示例/parse.go
 */
package main

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
    "github.com/dgrijalva/jwt-go"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        fmt.Print("No .env file found")
    }
}

func main() {
    tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMTIsImV4cCI6MTU3NTYzMzE2MCwiaXNzIjoibG9jYWxob3N0In0.tl0obbcF9me3LALR-3uCYuGSFcxK7kXRoFXGz4dj1fk"
    
    SecretString, exists := os.LookupEnv("SECRET_KEY")

    if !exists {
        return
    }

    fmt.Println(SecretString)
    Secret := []byte(SecretString)

    type MyCustomClaims struct {
        ID int `json:"id"`
        jwt.StandardClaims
    }
    
    token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return Secret, nil
    })

    if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
        fmt.Println(claims.ID, claims.StandardClaims.ExpiresAt)
    } else {
        fmt.Println(err)
    }
}