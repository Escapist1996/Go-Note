package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner在给定的超时时间内执行一组任务
type Runner struct {
	//interrupt通道报告从OS发送的信号
	interrupt chan os.Signal

	//complete通道报告处理任务已经完成
	complete chan error
	timeout <-chan time.Time

	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1)
		complete:  make(chan error)
		timeout:   time.After(d)
	}
}

//Add将一个任务附加到Runner上。这个任务是接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

//Start执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	//用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	//当任务处理完成时发出的信号
	case err := <-r.complete:
			  return err
	//当任务处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		
		//执行已注册任务
		task(id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return true

		//继续执行
		default:
			return false

	}
}


