package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func RouteHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is route 1")
}

func RouteHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is route 2")
}

func OtherRoutesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is for other routes")
}


func main() {
	// 创建路由器
	r := mux.NewRouter()

	// 添加路由处理函数
	r.HandleFunc("/route1", RouteHandler1)
	r.HandleFunc("/route2", RouteHandler2)

	// 配置其他路由走第三个配置
	r.PathPrefix("/").HandlerFunc(OtherRoutesHandler)

	// 启动服务器
	http.Handle("/", r)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)

}

/*
docker exec go_c go run demo/web.go
*/
