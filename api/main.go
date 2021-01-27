// main.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {

	router := httprouter.New()

	router.GET("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	
	return router

}
func main() {
	r := RegisterHandler()

	http.ListenAndServe(":8000", r)
}

