package session

import (
	"sync"
	"time"

	"github.com/gitcloneese/video_server/api/dbops"
	"github.com/gitcloneese/video_server/api/defs"
	"github.com/gitcloneese/video_server/api/utils"
)

var sessionMap *sync.Map

func nowMill() int64{
	return time.Now().UnixNano() / 1000000
}

func init() {
	sessionMap = &sync.Map{}

}

func LoadSessionsFromDB() {
	r, err := dbops.RetriveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})

}

// un username
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano() /1000000
	ttl := ct + 30 * 60 * 1000 //server side valid time :30 min
	
	ss := &defs.SimpleSession{Usemename : un, TTL :ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	
	return id

}

func DeleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
	
}


//sid sessions_id
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowMill()
		if ss.(*defs.SimpleSession).TTL < ct{
			//delete expired session
			DeleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Usemename, false
	}
	
	return "", true
	
}
