package pool

import "sync/atomic"

func (p *Pool) incRunningWorkers() {
	atomic.AddInt32(&p.runningWorkers, 1)
}

func (p *Pool) decRunningWorkers() {
	atomic.AddInt32(&p.runningWorkers, ^int32(0))
}

func (p *Pool) getRunningWorkers() int32 {
	return atomic.LoadInt32(&p.runningWorkers)
}

func (p *Pool) getCap() int32 {
	return atomic.LoadInt32(&p.capacity)
}
