package taskrunner

import (
	"errors"
	"log"

	"os"

	"sync"

	"../dbops"
	"time"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)

	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}

func VideoClearDisPatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All task finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil

}

func VideoClearExecutor(dc dataChan) error {

	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			// 闭包
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}

			}(vid)
		default:
			break forloop
		}
	}
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
		
	})
	time.Sleep(time.Second)

	return err
}
