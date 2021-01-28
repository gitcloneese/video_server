// main.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct{
	r *httprouter.Router
}

func  NewMiddleWareHandler(r *httprouter.Router) http.Handler{ //http.Handler 是个接口 实现了 ServerHTTP/2 方法
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}



func RegisterHandler() *httprouter.Router { //Router 也实现了ServeHTTP

	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	
	return router

}
func main() {
	r := RegisterHandler()
	mh := NewMiddleWareHandler(r)

	http.ListenAndServe(":8000", mh)
	
	// http.HandleFunc("/", RegisterHandler)
	// http.ListenAndServe(":8000", r)
}

