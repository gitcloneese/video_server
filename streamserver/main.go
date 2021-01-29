package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	
)

func RegisterHandler() *httprouter.Router{
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHander)
	router.POST("/upload:vid-id", uploadHandler)
	
	
	return router
}




func main(){
	r := RegisterHandler()
	http.ListenAndServe(":9000", r)
	
}