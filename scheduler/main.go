package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/gitcloneese/video_server/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router

}

func main() {
	taskrunner.S
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}
