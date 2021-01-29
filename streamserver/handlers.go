package main

import (
	"net/http"
	"os"
	// "io/ioutil"
	"github.com/julienschmidt/httprouter"
	"time"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl :=  VIDEO_DIR + vid    //video link
	video, err := os.Open(vl)	
	if err != nil{
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	
	//设置成视频mp4
	w.Header().Set("Content-Type", "video/mp4")
	
	http.ServeContent(w, r, "", time.Now(), video)   //参数3名字
	
	defer video.Close()

}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
