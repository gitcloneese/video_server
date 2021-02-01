package dbops

import (
	_	"github.com/go-sql-driver/mysql"
	"log"
	
)

func AddVideoDeletionRecord(vid string) error {
	stmtIn, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}
	
	_, err = stmtIn.Exec(vid)
	if err != nil{
		log.Println("AddVideoDeletionRecord error : %v", err)
		return err
	}
	
	defer stmtIn.Close()
	return nil
}