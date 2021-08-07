package main

import (
	"encoding/json"
	"os"

	"github.com/valyala/fasthttp"
)

func getSettings() (map[string]map[string]interface{}, error) {

	var settings map[string]map[string]interface{}

	rf, err := os.ReadFile("./settings.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rf, &settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func respError(ctx *fasthttp.RequestCtx, respErr ErrorResponse, statusCode int) error {
	rb, err := json.Marshal(respErr)

	if err != nil {
		return err
	}

	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.SetBody(rb)
	return nil
}

func respSuccess(ctx *fasthttp.RequestCtx, respSuccess SuccessResponse, statusCode int) error {
	rb, err := json.Marshal(respSuccess)

	if err != nil {
		return err
	}

	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.SetBody(rb)
	return nil
}
