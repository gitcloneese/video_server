package main

import (
	"io"
	"net/http"
	"github.com/gitcloneese/video_server/api/defs"
	"encoding/json"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)
	
	resStr, _ := json.Marshal(&errResp.Error)
	
	io.WriteString(w, string(resStr))

}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
