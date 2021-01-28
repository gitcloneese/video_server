package dbops

import (
	"strconv"
	"sync"

	"database/sql"

	"log"

	"github.com/gitcloneese/video_server/api/defs"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare(`INSERT INTO sessions (session_id, TTL, login_name) 
		VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetriveSession(sid string) (*defs.SimpleSession, error) {

	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare(`SELECT TTL, login_name FROM sessions WHERE 
		session_id = ?`)

	if err != nil {
		return nil, err
	}

	var ttl, uname string

	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err != nil {
		return nil, err
	} else {
		ss.TTL = res
		ss.Usemename = uname
	}

	defer stmtOut.Close()
	return ss, nil
}

func RetriveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}

	stmtOut, err := dbConn.Prepare(`SELECT * FROM sessions`)

	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}

	defer stmtOut.Close()

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string

		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrieve all sessions error : %v\n", err)
			continue
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Usemename: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id : %s, ttl: %d\n", id, ss.TTL)
		} else {
			log.Printf("parse TTL error: %s \n", err1)
			continue
		}
	}
	return m, nil

}

func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare(`DELETE FROM sessions WHERE 
	session_id = ?`)
	if err != nil {
		return err
	}

	if _, err := stmtDel.Exec(sid); err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
