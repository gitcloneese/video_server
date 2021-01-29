package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"html/template"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid //video link
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Open file error : %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	//设置成视频mp4
	w.Header().Set("Content-Type", "video/mp4")

	http.ServeContent(w, r, "", time.Now(), video) //参数3名字

	defer video.Close()

}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file") //
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server Error")
		return
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Printf("Read file error : %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server Error")
	}

	filename := p.ByName("vid-id")

	path := VIDEO_DIR + filename

	err = ioutil.WriteFile(path, data, 0666) //

	if err != nil {
		log.Printf("write file error : %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload successfully")

}

func TestPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}
