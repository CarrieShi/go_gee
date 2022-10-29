package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

/////////// 注册 仅过程化  ////////
func SignUpProgress(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "deserialized failed: %v", err)
		return
	}

	// 返回虚拟的 userId，假装注册成功了
	resp := &commonResponse{
		Data: rand.Int(),
	}
	respJson, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(w, "resp serialized failed: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "%d", rand.Int())
	// respJson byte 切片可转string
	fmt.Fprintf(w, string(respJson))

}
