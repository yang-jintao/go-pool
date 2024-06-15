package pool

import (
	"errors"
	"fmt"
	"log"
)

const (
	Running = 1
	Stop    = 2
)

var poolAlreadyClosedErr = errors.New("pool has closed")
var invalidPoolCapErr = errors.New("invalid cap")

type Task struct {
	Handler func()
}

type Pool struct {
	capacity       int32
	runningWorkers int32
	state          int32
	taskChannel    chan *Task
	closeChannel   chan bool
	PanicHandler   func(interface{})
}

func NewPool(size int32) (*Pool, error) {
	if size <= 0 {
		return nil, invalidPoolCapErr
	}

	return &Pool{
		capacity:       size,
		runningWorkers: 0,
		state:          Running,
		taskChannel:    make(chan *Task, size),
		closeChannel:   make(chan bool),
	}, nil
}

func (p *Pool) Put(task *Task) error {
	if p.state == Stop {
		return poolAlreadyClosedErr
	}

	if p.getRunningWorkers() < p.getCap() {
		p.Run()
	}
	p.taskChannel <- task // 这里如果满了是阻塞

	return nil
}

func (p *Pool) Run() {
	p.incRunningWorkers()

	go func() {
		defer func() {
			p.decRunningWorkers()
			if err := recover(); err != nil {
				if p.PanicHandler != nil {
					p.PanicHandler(err)
				} else {
					log.Printf("Panic Err=%+v\n", err)
				}
			}
		}()

		for {
			select {
			case task, ok := <-p.taskChannel:
				if !ok {
					fmt.Println("not ok ")
					return
				}
				task.Handler()
				//case <-p.closeChannel:
				//	fmt.Println("has closed ")
				//	return
			}
		}
	}()
}

func (p *Pool) Close() {
	if p.state == Stop {
		return
	}

	p.state = Stop
	for len(p.taskChannel) > 0 {
	}

	//p.closeChannel <- true
	close(p.taskChannel)
}
