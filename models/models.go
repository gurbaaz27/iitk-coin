package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Name     string `json:"name"`
	Rollno   int64  `json:"rollno,string"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Rollno   int64  `json:"rollno,string"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Rollno int64 `json:"rollno,string"`
	jwt.StandardClaims
}

type LoginToken struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type RewardPayload struct {
	Rollno int64 `json:"rollno,string"`
	Coins  int64 `json:"coins,string"`
}

type TransferPayload struct {
	SenderRollno   int64 `json:"sender,string"`
	ReceiverRollno int64 `json:"receiver,string"`
	Coins          int64 `json:"coins,string"`
}
