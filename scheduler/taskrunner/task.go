package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/gitcloneese/video_server/scheduler/dbops"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	//删除出错， 且错误不是文件不存在
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}

func VideoClearDisPatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error :%v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {

	errMap := &sync.Map{}

forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(vid.(string)); err != nil {
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
		err := v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return nil

}
