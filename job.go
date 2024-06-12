package threadPool

type Job interface {
	RunTask(request interface{}) //用户定制化的任务
}

type JobChan chan Job //高并发需要执行的任务
