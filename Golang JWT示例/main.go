/*
 * @Author: your name
 * @Date: 2021-04-08 14:55:49
 * @LastEditTime: 2021-04-08 15:40:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/Golang JWT示例/main.go
 */
package main

import (
    "os"
    "fmt"
    "time"
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
    SecretString, exists := os.LookupEnv("SECRET_KEY")

    if !exists {
        return
    }

    fmt.Println(SecretString)
    Secret := []byte(SecretString) // 标准的glang显式转换, 将变量SecretString转换成[]byte类型, 并且赋值给Secret

    type MyCustomClaims struct {
        ID int `json:"id"`
        jwt.StandardClaims
    }
    
    // Create the Claims
    claims := MyCustomClaims{
        10012,
        jwt.StandardClaims{
            ExpiresAt: time.Time.AddDate(time.Now(),0, 0, 7).Unix(),
            Issuer:    "localhost",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    ss, err := token.SignedString(Secret)
    fmt.Printf("%v %v", ss, err)
}