package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/// 中间件 流控
type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWardHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.l.GetConn() == false {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too Many Requests")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

/// --流控

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", StreamHandler)
	router.POST("/upload/:vid-id", UploadHandler)
	router.GET("/testpage", TestPageHandler)

	return router
}

func main() {
	r := RegisterHandler()
	mh := NewMiddleWardHandler(r, 2)
	http.ListenAndServe(":9000", mh)

}
