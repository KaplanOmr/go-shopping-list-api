package main

import (
	"strings"

	"github.com/valyala/fasthttp"
)

func rootHandlers(ctx *fasthttp.RequestCtx) {
	route := string(ctx.Request.RequestURI())
	method := string(ctx.Request.Header.Method())
	as := strings.Split(string(ctx.Request.Header.Peek("Authorization")), " ")

	if len(as) != 2 {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 10004
		respErr.ErrorMsg = "AUTHORIZATION_INVALID"

		respError(ctx, respErr, 401)
		return
	}

	token := as[1]

	authAllowedURI := []string{
		"/",
		"/register",
		"/login",
	}

	if !authCheck(ctx, token, route, authAllowedURI) {
		return
	}

	switch route {
	case "/":
		mainHandler(ctx, method)
	case "/register":
		registerHandler(ctx, method)
	case "/login":
		loginHandler(ctx, method)
	default:
		var resp ErrorResponse

		resp.Status = false
		resp.ErrorCode = 10001
		resp.ErrorMsg = "INVALID_URI"

		respError(ctx, resp, 400)
	}
}

func mainHandler(ctx *fasthttp.RequestCtx, method string) {

	if !allowedMethod(ctx, method, "GET") {
		return
	}

	var resp SuccessResponse

	resp.Status = true
	resp.Data = "thanx"

	respSuccess(ctx, resp, 200)
}

func registerHandler(ctx *fasthttp.RequestCtx, method string) {
	if !allowedMethod(ctx, method, "POST") {
		return
	}

	params := []string{
		"username",
		"email",
		"name",
		"password",
	}

	if !reqParams(ctx, params) {
		return
	}

	var adb AppDB

	adb.conDB()
	defer adb.DB.Close()

	var u UserStruct

	u.Email = string(ctx.PostArgs().Peek("email"))
	u.Username = string(ctx.PostArgs().Peek("username"))
	u.Name = string(ctx.PostArgs().Peek("name"))
	u.Password = string(ctx.PostArgs().Peek("password"))

	if ec, check := adb.createUser(u); !check {
		var respErr ErrorResponse

		respErr.Status = false

		switch ec {
		case 1:
			respErr.ErrorCode = 100011
			respErr.ErrorMsg = "REGISTER_SERVER_ERROR"
		case 2:
			respErr.ErrorCode = 100012
			respErr.ErrorMsg = "REGISTER_USERNAME_INVALID"
		case 3:
			respErr.ErrorCode = 100013
			respErr.ErrorMsg = "REGISTER_EMAIL_INVALID"
		}

		respError(ctx, respErr, 400)
		return
	}

	var resp SuccessResponse

	resp.Status = true
	resp.Data = "CREATED_USER"

	respSuccess(ctx, resp, 200)
}

func loginHandler(ctx *fasthttp.RequestCtx, method string) {
	if !allowedMethod(ctx, method, "POST") {
		return
	}

	params := []string{
		"username",
		"password",
	}

	if !reqParams(ctx, params) {
		return
	}

	var adb AppDB

	adb.conDB()
	defer adb.DB.Close()

	var u UserStruct

	u.Password = string(ctx.PostArgs().Peek("password"))
	u.Username = string(ctx.PostArgs().Peek("username"))

	ru, check := adb.checkUser(u)

	if !check {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100020
		respErr.ErrorMsg = "USER_INFO_INCORRET"

		respError(ctx, respErr, 400)
		return
	}

	t, err := createToken(ru)

	if err != nil {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100021
		respErr.ErrorMsg = "LOGIN_ERROR"

		respError(ctx, respErr, 400)
		return
	}

	var resp SuccessResponse

	resp.Status = true
	resp.Data = t

	respSuccess(ctx, resp, 200)

}
