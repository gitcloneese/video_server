package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"github.com/gitcloneese/video_server/api/defs"
	"encoding/json"
	"github.com/gitcloneese/video_server/api/dbops"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	
	if  err := json.Unmarshal(res, ubody); err != nil{
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	
	if err = dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil{
		sendErrorResponse(w, defs.ErrorDBError)
	}
	
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	uname := p.ByName("user_name")	
	io.WriteString(w, uname)
	
}