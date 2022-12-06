package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
	从 http.Request 中读取数据并解析
	从 http.ResponseWriter 中写入数据和相应
	json数据的序列化和反序列化
*/

func SignUpWithoutContext(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed", err)
		return
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "unmarshal failed", err)
		return
	}
	fmt.Fprintf(w, "%d", err)
}

func main() {
}

type signUpReq struct {
	Email             string `json:"email`
	Password          string `json:password`
	ConfirmedPassword string `json:confirmed_password`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
