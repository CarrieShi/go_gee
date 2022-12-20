package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// 用户创建context ，not commend ，see signUp
// 希望上下文依赖框架创建，因为如果有其他的 interface, 用户可能会不知道使用哪个
func SignUpByCustomContext(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}

	ctx := &Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	resp := &commonResponse{
		Data: rand.Int(),
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	//ctx.W.Write(err.Error())//写入都可能失败了，该记录日志了
	if err != nil {
		fmt.Fprintf(w, "写入响应失败: %v", err)
		return
	}
}

func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: rand.Int(),
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		ctx.SystemErrorJson(err)
		return
	}
}
