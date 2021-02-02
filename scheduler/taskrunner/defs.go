package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTOR = "e"
	CLOSE             = "c"
	
	VIDEO_PATH  = "./videos/"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
