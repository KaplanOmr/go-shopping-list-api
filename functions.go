package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/valyala/fasthttp"
)

func getSettings() map[string]map[string]interface{} {

	var s map[string]map[string]interface{}

	rf, err := os.ReadFile("./settings.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(rf, &s)
	if err != nil {
		panic(err)
	}

	return s
}

func respError(ctx *fasthttp.RequestCtx, respErr ErrorResponse, statusCode int) {

	rb, err := json.Marshal(respErr)

	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"# RESPONSE: %s | STATUS: false | BODY: %s \n",
		ctx.RemoteAddr().String(),
		string(rb),
	)

	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.SetBody(rb)
}

func respSuccess(ctx *fasthttp.RequestCtx, respSuccess SuccessResponse, statusCode int) {
	rb, err := json.Marshal(respSuccess)

	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"# RESPONSE: %s | STATUS: true | BODY: %s \n",
		ctx.RemoteAddr().String(),
		string(rb),
	)

	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.SetBody(rb)
}

func allowedMethod(ctx *fasthttp.RequestCtx, in string, allow string) bool {
	if in != allow {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 10002
		respErr.ErrorMsg = "INVALID_REQUEST_METHOD"

		respError(ctx, respErr, 400)

		return false
	}

	return true
}

func reqPostParams(ctx *fasthttp.RequestCtx, params []string) bool {

	c := true

	for _, p := range params {
		if !ctx.PostArgs().Has(p) {
			c = false
		}
	}

	if !c {
		var resp ErrorResponse

		resp.Status = false
		resp.ErrorCode = 10010
		resp.ErrorMsg = "REQUIRED_PARAMS_INVALID"
		resp.Data = map[string][]string{
			"required_params": params,
		}

		respError(ctx, resp, 200)

	}

	return c
}

func reqGetParams(ctx *fasthttp.RequestCtx, params []string) bool {

	c := true

	for _, p := range params {
		if !ctx.QueryArgs().Has(p) {
			c = false
		}
	}

	if !c {
		var resp ErrorResponse

		resp.Status = false
		resp.ErrorCode = 10010
		resp.ErrorMsg = "REQUIRED_PARAMS_INVALID"
		resp.Data = map[string][]string{
			"required_params": params,
		}

		respError(ctx, resp, 200)

	}

	return c
}

func createToken(u UserStruct) (string, error) {

	var signPri []byte

	signPri, err := getKeys("private")

	if err != nil {
		return "", err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signPri)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user": u,
	})

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func checkToken(tokenString string) (UserStruct, bool) {

	var signPub []byte

	signPub, err := getKeys("public")

	if err != nil {
		fmt.Println(err)
		return UserStruct{}, false
	}

	signKey, err := jwt.ParseRSAPublicKeyFromPEM(signPub)

	if err != nil {
		fmt.Println(err)
		return UserStruct{}, false
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return signKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return UserStruct{}, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uraw := claims["user"].(map[string]interface{})

		u := UserStruct{
			ID:       int(uraw["id"].(float64)),
			Name:     uraw["name"].(string),
			Email:    uraw["email"].(string),
			Username: uraw["username"].(string),
		}

		return u, true
	} else {
		return UserStruct{}, false
	}
}

func getKeys(keyType string) ([]byte, error) {
	var readedFile string

	if keyType == "private" {
		readedFile = settings["jwt"]["pri"].(string)
	} else if keyType == "public" {
		readedFile = settings["jwt"]["pub"].(string)
	} else {
		return nil, fmt.Errorf("invalid key type: %s", keyType)
	}

	rf, err := os.ReadFile(readedFile)

	if err != nil {
		return nil, err
	}

	return rf, nil
}

func authCheck(ctx *fasthttp.RequestCtx, auth []string, uri string, allow []string) (UserStruct, bool) {

	for _, val := range allow {
		if val == uri {
			return UserStruct{}, true
		}
	}

	if len(auth) != 2 {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 10004
		respErr.ErrorMsg = "AUTHORIZATION_INVALID"

		respError(ctx, respErr, 401)
		return UserStruct{}, false
	}

	token := auth[1]

	u, c := checkToken(token)

	if !c {
		var respErr ErrorResponse

		respErr.Status = false
		respErr.ErrorCode = 10003
		respErr.ErrorMsg = "TOKEN_INVALID"

		respError(ctx, respErr, 401)
		return UserStruct{}, false
	}

	return u, true
}
