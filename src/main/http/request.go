package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	server := NewHttpServer("test-server")
	server.RouteRW("/user/SignUpProgress", SignUpProgress)
	server.Route("/user/signUp", SignUp)
	server.Start(":8080")

	//http.HandleFunc("/body/once", readBodyOnce)
	//http.HandleFunc("/body/multi", getBodyNil)
	//http.HandleFunc("/url/query", getParams)
	//http.HandleFunc("/url/whole", wholeUrl)
	//http.HandleFunc("/header", header)
	//http.HandleFunc("/form", form)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "after parse form %v\n", r.Form)
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v\n", r.Header)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

func getParams(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Fprintf(w, "query is %v \n", query)
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
