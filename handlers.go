package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func rootHandlers(ctx *fasthttp.RequestCtx) {
	route := string(ctx.Request.RequestURI())
	method := string(ctx.Request.Header.Method())

	fmt.Println(string(ctx.Request.Header.Peek("Authorization")))

	switch route {
	case "/":
		allowedMethod(ctx, method, "GET")
		mainHandler(ctx, method)
	case "/register":
		allowedMethod(ctx, method, "POST")
	default:
		var resp ErrorResponse

		resp.Status = false
		resp.ErrorCode = 10001
		resp.ErrorMsg = "INVALID_URI"

		respError(ctx, resp, 400)
	}
}

func mainHandler(ctx *fasthttp.RequestCtx, method string) {

	if method != "GET" {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 10002
		respErr.ErrorMsg = "INVALID_REQUEST_METHOD"

		respError(ctx, respErr, 400)
		return
	}

	var resp SuccessResponse

	resp.Status = true
	resp.Data = "thanx"

	respSuccess(ctx, resp, 200)
}

func registerHandler(ctx *fasthttp.RequestCtx, method string) {

}
