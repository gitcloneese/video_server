package taskrunner

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	datasize   int
	longlived  bool
	Dispather  fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longlived:  longlived,
		datasize:   size,
		Dispather:  d, //生产者
		Executor:   e, //消费者
	}
}

func (r *Runner) startDisPatch() {
	defer func() {
		if r.longlived == false {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}

	}()
	for {
		select {
		case c := <-r.Controller:
			if c == REDAY_TO_DISPATCH {
				err := r.Dispather(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- REDAY_TO_DISPATCH
				}
			}

		case e := <-r.Error:
			if e == CLOSE {
				return
			}

			// default :
		}
	}
}

func (r *Runner) StartAll() {
	/*
		开启前要放一个 REDAY_TO_DISPATCH ,要不然永远都进不去 
		处理data
	*/
	r.Controller <- REDAY_TO_DISPATCH
	r.startDisPatch()

}
