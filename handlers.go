package main

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

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

	if !reqPostParams(ctx, params) {
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

	if !reqPostParams(ctx, params) {
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

func listCreateHandler(ctx *fasthttp.RequestCtx, method string, userData UserStruct) {
	if !allowedMethod(ctx, method, "POST") {
		return
	}

	params := []string{
		"title",
		"total_cost",
		"status",
	}

	if !reqPostParams(ctx, params) {
		return
	}

	tc, err := strconv.ParseFloat(string(ctx.PostArgs().Peek("total_cost")), 64)

	if err != nil {
		fmt.Println(err)
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100044
		respErr.ErrorMsg = "TOTAL_COST_PARAMETER_INVALID"

		respError(ctx, respErr, 400)
		return
	}

	status, err := strconv.Atoi(string(ctx.PostArgs().Peek("status")))

	if err != nil {
		fmt.Println(err)
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100044
		respErr.ErrorMsg = "STATUS_PARAMETER_INVALID"

		respError(ctx, respErr, 400)
		return
	}

	var newList ListStruct

	newList.Title = string(ctx.PostArgs().Peek("title"))
	newList.TotalCost = tc
	newList.Status = status

	var adb AppDB

	adb.conDB()
	defer adb.DB.Close()

	l, c := adb.listCreate(userData, newList)

	if !c {
		fmt.Println(err)
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100084
		respErr.ErrorMsg = "LIST_CREATE_ERR"

		respError(ctx, respErr, 500)
		return
	}

	var resp SuccessResponse

	resp.Status = true
	resp.Data = l

	respSuccess(ctx, resp, 200)
}

func listGetAllHandler(ctx *fasthttp.RequestCtx, method string, userData UserStruct) {
	if !allowedMethod(ctx, method, "POST") {
		return
	}

	params := []string{
		"sortby",
		"limit",
		"offset",
	}

	if !reqPostParams(ctx, params) {
		return
	}

	desc := string(ctx.PostArgs().Peek("sortby"))
	limit := string(ctx.PostArgs().Peek("limit"))
	offset := string(ctx.PostArgs().Peek("offset"))

	var adb AppDB

	adb.conDB()
	defer adb.DB.Close()

	lists, c := adb.listGetAll(userData, desc, limit, offset)

	if !c {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100085
		respErr.ErrorMsg = "LIST_GET_ERR"

		respError(ctx, respErr, 500)
		return
	}

	var resp SuccessResponse

	resp.Status = true
	resp.Data = lists

	respSuccess(ctx, resp, 200)
}

func listGetHandler(ctx *fasthttp.RequestCtx, method string, userData UserStruct) {
	if !allowedMethod(ctx, method, "GET") {
		return
	}

	params := []string{"id"}

	if !reqGetParams(ctx, params) {
		return
	}

	id, err := strconv.Atoi(string(ctx.QueryArgs().Peek("id")))

	if err != nil {
		fmt.Println(err)
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100086
		respErr.ErrorMsg = "ID_PARAMETER_INVALID"

		respError(ctx, respErr, 400)
		return
	}

	var adb AppDB
	adb.conDB()
	defer adb.DB.Close()

	l, c := adb.listGet(userData, id)

	if !c {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100085
		respErr.ErrorMsg = "LIST_GET_ERR"

		respError(ctx, respErr, 500)
		return
	}

	var resp SuccessResponse

	resp.Status = true
	resp.Data = l

	respSuccess(ctx, resp, 200)
}

func listUpdateHandler(ctx *fasthttp.RequestCtx, method string, userData UserStruct) {
	if !allowedMethod(ctx, method, "POST") {
		return
	}

	getParams := []string{"id"}
	postParams := []string{"title", "status"}

	if !reqGetParams(ctx, getParams) || !reqPostParams(ctx, postParams) {
		return
	}

	id, err := strconv.Atoi(string(ctx.QueryArgs().Peek("id")))

	if err != nil {
		fmt.Println(err)
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100086
		respErr.ErrorMsg = "ID_PARAMETER_INVALID"

		respError(ctx, respErr, 400)
		return
	}

	title := string(ctx.PostArgs().Peek("title"))
	status, err := strconv.Atoi(string(ctx.PostArgs().Peek("status")))

	if err != nil {
		fmt.Println(err)
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100086
		respErr.ErrorMsg = "STATUS_PARAMETER_INVALID"

		respError(ctx, respErr, 400)
		return
	}

	var adb AppDB
	adb.conDB()
	defer adb.DB.Close()

	c := adb.listUpdate(userData, title, status, id)

	if !c {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100095
		respErr.ErrorMsg = "LIST_UP_ERR"

		respError(ctx, respErr, 500)
		return
	}

	var resp SuccessResponse

	resp.Status = true

	respSuccess(ctx, resp, 200)
}

func listDeleteHandler(ctx *fasthttp.RequestCtx, method string, userData UserStruct) {
	if !allowedMethod(ctx, method, "GET") {
		return
	}

	params := []string{"id"}

	if !reqGetParams(ctx, params) {
		return
	}

	id, err := strconv.Atoi(string(ctx.QueryArgs().Peek("id")))

	if err != nil {
		fmt.Println(err)
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100086
		respErr.ErrorMsg = "ID_PARAMETER_INVALID"

		respError(ctx, respErr, 400)
		return
	}

	var adb AppDB
	adb.conDB()
	defer adb.DB.Close()

	c := adb.listDelete(userData, id)

	if !c {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 100099
		respErr.ErrorMsg = "LIST_DEL_ERR"

		respError(ctx, respErr, 500)
		return
	}

	var resp SuccessResponse

	resp.Status = true

	respSuccess(ctx, resp, 200)
}
