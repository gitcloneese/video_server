package main

import (
	"net/http"

	"github.com/gitcloneese/video_server/scheduler/taskrunner"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video_delete-rec/:vid-id", vidDelRecHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
