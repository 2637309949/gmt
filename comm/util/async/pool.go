package async

import (
	"sync"
	"sync/atomic"
)

const (
	start int32 = iota
	running
	shutdown
	stop
)

type simpleFSM struct {
	status int32
}

func newSimpleFSM() *simpleFSM {
	return &simpleFSM{
		status: start,
	}
}

func (s *simpleFSM) actEvent(stat int32) bool {
	if s.Current() > stat {
		return false
	} else {
		return atomic.CompareAndSwapInt32(&s.status, s.status, stat)
	}
}

func (s *simpleFSM) isRunning() bool {
	return s.Current() == running
}

func (s *simpleFSM) Current() int32 {
	return atomic.LoadInt32(&s.status)
}

func (s *simpleFSM) Action(stat int32) bool {
	switch stat {
	case start:
	case running:
	case shutdown:
	case stop:
		return s.actEvent(stat)
	default:
		return false
	}
	return false
}

type basePoolExecutor struct {
	corePoolSize chan int
	fsm          *simpleFSM
	oc           sync.Once
}

type PoolExecutor struct {
	basePoolExecutor
}

type WaitPoolExecutor struct {
	basePoolExecutor
	gp *sync.WaitGroup
}

func newBasePoolExecutor(cap int) basePoolExecutor {
	return basePoolExecutor{
		corePoolSize: make(chan int, cap),
		fsm:          newSimpleFSM(),
		oc:           sync.Once{},
	}
}

func NewPool(cap int) *PoolExecutor {
	checkCap(cap)
	return &PoolExecutor{
		basePoolExecutor: newBasePoolExecutor(cap),
	}
}

func NewWaitPool(cap int) *WaitPoolExecutor {
	checkCap(cap)
	return &WaitPoolExecutor{
		basePoolExecutor: newBasePoolExecutor(cap),
		gp:               new(sync.WaitGroup),
	}
}

func checkCap(cap int) {
	if cap < 0 {
		panic("The pool cap cannot lower zero")
	}
}

// checkSubmit defined TODO
func (b *basePoolExecutor) checkSubmit(f func()) {
	if f == nil {
		panic("The submit func is nil")
	}
	b.oc.Do(func() {
		b.fsm.actEvent(running)
	})
	if !b.fsm.isRunning() {
		panic("The pool is not running")
	}
}

// ShutDown defined TODO
func (b *basePoolExecutor) ShutDown() {
	b.fsm.actEvent(shutdown)
}

// IsShutDown defined TODO
func (b *basePoolExecutor) IsShutDown() bool {
	return b.fsm.Current() >= shutdown
}

// IsTerminated defined TODO
func (b *basePoolExecutor) IsTerminated() bool {
	return b.fsm.Current() >= stop
}

// Submit defined TODO
func (t *PoolExecutor) Submit(f func()) {
	t.checkSubmit(f)
	t.corePoolSize <- 1
	go func() {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
			<-t.corePoolSize
		}()
		if t.IsTerminated() {
			return
		}
		f()
	}()
}

// Submit defined TODO
func (t *WaitPoolExecutor) Submit(f func()) {
	t.checkSubmit(f)
	t.gp.Add(1)
	t.corePoolSize <- 1
	go func() {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
			<-t.corePoolSize
			t.gp.Done()
		}()
		if t.IsTerminated() {
			return
		}
		f()
	}()
}

// Wait defined TODO
func (t *WaitPoolExecutor) Wait() {
	t.gp.Wait()
	t.fsm.actEvent(stop)
	close(t.corePoolSize)
}
