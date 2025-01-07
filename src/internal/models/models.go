package models

import "github.com/golang-jwt/jwt/v5"

type Status struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Points int    `json:"points"`
}

type Referrer struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

type User struct {
	ID           int     `json:"id"`
	TelegramID   int64   `json:"telegram_id"`
	Active       bool    `json:"active"`
	RegisteredAt string  `json:"registered_at"`
	UserName     string  `json:"user_name"`
	StatusID     int     `json:"status_id"`
	LastVisit    *string `json:"last_visit"`
	Balance      int64   `json:"balance"`
	IsAdmin      bool    `json:"isadmin"`
	ReferrerID   *int    `json:"referrer_id"`
}

type ResponseUser struct {
	ID           int     `json:"id"`
	TelegramID   int64   `json:"telegram_id"`
	Active       bool    `json:"active"`
	RegisteredAt string  `json:"registered_at"`
	UserName     string  `json:"user_name"`
	StatusID     int     `json:"status_id"`
	LastVisit    *string `json:"last_visit"`
	Balance      int64   `json:"balance"`
	IsAdmin      bool    `json:"isadmin"`
	ReferrerID   *int    `json:"referrer_id"`
	Cache        bool    `json:"cache"`
}

type TaskActivity struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	TaskID int `json:"task_id"`
}

type PostResponse struct {
	Task string `json:"task"`
}

type ErrorResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

type PostReferrer struct {
	Code string `json:"code"`
}

type JwtCustomClaims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
