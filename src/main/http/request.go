package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/multi", getBodyNil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBodyNil(w http.ResponseWriter, r *http.Request) {
	//body, _ := r.GetBody()
	//readbody, _ := io.ReadAll(body)
	//fmt.Fprintf(w, "read the data 1 : %s \n", string(readbody))

	//body, _ = r.GetBody()
	//readbody, _ = io.ReadAll(body)
	//
	//fmt.Fprintf(w, "read the data 2 : %s \n", string(readbody))

	if r.GetBody == nil {
		// 原生 进入这里了
		fmt.Fprint(w, "get body is nil \n")
	} else {
		fmt.Fprint(w, "get body is not nil \n")
	}

}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}

	fmt.Fprintf(w, "read the data : %s \n", string(body))

	// 读第二次，读不到，不报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不进来
		fmt.Fprintf(w, "read body failed twice: %v", err)
		return
	}

	fmt.Fprintf(w, "read the data twice : %s \n", string(body))

	// 读完，不会 reset body 内容，即 body 已经为空
}
