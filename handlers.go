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

	if !authCheck(ctx, token, route, []string{"/", "/register"}) {
		return
	}

	switch route {
	case "/":
		mainHandler(ctx, method)
	case "/register":
		registerHandler(ctx, method)
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
		"ok",
		"ok2",
		"ok3",
	}

	if !reqParams(ctx, params) {
		return
	}
}
