package threadPool

//线程

type Worker struct {
	//不需要带缓冲的任务队列
	JobQueue JobChan
	//退出标志
	Quit chan bool
}

//创建一个新的线程对象

func NewWorker() Worker {
	return Worker{
		make(JobChan),
		make(chan bool),
	}
}

//启动一个线程,来监听Job事件
//执行完任务,需要将自己重新发送到workerPool

func (w Worker) Start(workerPool *WorkerPool) {
	go func() {
		for {
			//将线程注册到线程池
			workerPool.WorkerQueue <- &w
			select {
			case job := <-w.JobQueue:
				job.RunTask(nil)
			//终止当前线程
			case <-w.Quit:
				return
			}
		}
	}()
}
