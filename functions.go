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

func createToken() (string, error) {

	var signPri []byte

	signPri, err := getKeys("private")

	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": UserStruct{
			ID:       1,
			Username: "omer",
		},
	})

	tokenString, err := token.SignedString(signPri)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func checkToken(tokenString string) bool {

	var signPub []byte

	signPub, err := getKeys("public")

	if err != nil {
		panic(err)
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return signPub, nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userData = claims["user"]
		return true
	} else {
		return false
	}
}

func getKeys(keyType string) ([]byte, error) {
	var readedFile string

	if keyType == "private" {
		readedFile = settings["jwt"]["pri"].(string)
	} else if keyType == "public" {
		readedFile = settings["jwt"]["pri"].(string)
	} else {
		return nil, fmt.Errorf("invalid key type: %s", keyType)
	}

	rf, err := os.ReadFile(readedFile)

	if err != nil {
		return nil, err
	}

	return rf, nil
}
