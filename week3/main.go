package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"time"
)

func main() {
	http.HandleFunc("/test", handler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("http server start err:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k,v := range r.Header {
		for i := 0; i < len(v); i++ {
			w.Header().Set(k, v[i])
		}
	}
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	w.WriteHeader(http.StatusOK)
	responseStatusCode := reflect.ValueOf(w).Elem().FieldByName("status").Int()
	fmt.Println(fmt.Sprintf("[log]请求时间:%s, 客户端IP:%s, HTTP返回码:%d ",
		time.Now().Format("2006-01-02 15:04:05"),
		r.RemoteAddr,
		responseStatusCode))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}
