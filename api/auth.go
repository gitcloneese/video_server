package main

import (
	"net/http"
	"github.com/gitcloneese/video_server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool{
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	
	uname, ok := session.IsSessionExpired(sid)	
	
	
}