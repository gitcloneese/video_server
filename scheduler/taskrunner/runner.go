package taskrunner

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	datasize   int
	longlived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longlived:  longlived,
		datasize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatcher() {
	defer func() {
		if !r.longlived {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}

	}()
	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE

				} else {
					r.Controller <- READY_TO_EXECUTOR
				}
			}
			
			if c == READY_TO_EXECUTOR {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}

		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}

}


func (r *Runner) startAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatcher()
}
