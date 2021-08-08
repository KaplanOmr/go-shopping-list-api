package main

var (
	userData interface{}
	settings = getSettings()
)

type ErrorResponse struct {
	Status    bool        `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
	Data      interface{} `json:"data"`
}

type SuccessResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ListStruct struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	UserID    int    `json:"user_id"`
	Status    int    `json:"status"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type UserStruct struct {
	ID       int
	Username string
	Password string
	Name     string
	Email    string
}
