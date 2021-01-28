package session

import (
	"time"
	"sync"
	"github.com/gitcloneese/video_server/api/defs"
)



var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
	
	
}

func LoadSessionsFromDB(){
	
}
// un username
func GenerateNewSessionId(un string)string{
	
}

//sid sessions_id
func IsSessionExpired(sid string)(string, bool){
	
}

