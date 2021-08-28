package main

import (
	"strings"

	"github.com/valyala/fasthttp"
)

func rootHandlers(ctx *fasthttp.RequestCtx) {
	route := string(ctx.Request.RequestURI())
	method := string(ctx.Request.Header.Method())
	auth := strings.Split(string(ctx.Request.Header.Peek("Authorization")), " ")

	authAllowedURI := []string{
		"/",
		"/register",
		"/login",
	}

	userData, authCheck := authCheck(ctx, auth, route, authAllowedURI)

	if !authCheck {
		return
	}

	switch route {
	case "/":
		mainHandler(ctx, method)
	case "/register":
		registerHandler(ctx, method)
	case "/login":
		loginHandler(ctx, method)
	case "/list/create":
		listCreateHandler(ctx, method, userData)
	case "/list/get_all":
		listGetAllHandler(ctx, method, userData)
	default:
		var resp ErrorResponse

		resp.Status = false
		resp.ErrorCode = 10001
		resp.ErrorMsg = "INVALID_URI"

		respError(ctx, resp, 400)
	}
}
