package main

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

func rootHandlers(ctx *fasthttp.RequestCtx) {
	fmt.Printf(
		"# REQUEST: %s | HEADER: %s | GET: %s | BODY: %s \n",
		ctx.RemoteAddr().String(),
		string(ctx.Request.Header.Peek("Authorization")),
		string(ctx.RequestURI()),
		string(ctx.PostArgs().QueryString()),
	)

	route := string(ctx.Request.URI().Path())
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
	case "/list/get":
		listGetHandler(ctx, method, userData)
	case "/list/update":
		listUpdateHandler(ctx, method, userData)
	case "/list/delete":
		listDeleteHandler(ctx, method, userData)
	default:
		var resp ErrorResponse

		resp.Status = false
		resp.ErrorCode = 10001
		resp.ErrorMsg = "INVALID_URI"

		respError(ctx, resp, 400)
	}
}
