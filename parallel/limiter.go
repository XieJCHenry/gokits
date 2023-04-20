package parallel

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type QueueStrategy string

const (
	Accepted QueueStrategy = "accepted"
	Rejected QueueStrategy = "rejected"
)

type LimiterJob struct {
	Params   []interface{}
	Executor func(params ...interface{}) LimiterJobResult
}

type executeJob struct {
	wg  *sync.WaitGroup
	job *LimiterJob
}

type LimiterJobResult struct {
	Data  interface{}
	Error error
}

type Limiter struct {
	strategy    QueueStrategy
	wg          sync.WaitGroup
	mtx         sync.Mutex
	started     atomic.Bool
	jobInChan   chan *executeJob
	jobDownChan chan struct{}
	buffer      []*executeJob
}

func NewLimiter(limit int, strategy QueueStrategy) *Limiter {
	if limit <= 0 {
		panic(fmt.Sprintf("limitation must be positive, not %d", limit))
	}
	if strategy != Accepted && strategy != Rejected {
		panic(fmt.Sprintf("limiter startegy must be either %s or %s", Accepted, Rejected))
	}

	limiter := new(Limiter)
	limiter.strategy = strategy
	limiter.wg = sync.WaitGroup{}
	limiter.mtx = sync.Mutex{}
	limiter.started = atomic.Bool{}
	limiter.jobInChan = make(chan *executeJob, limit)
	limiter.jobDownChan = make(chan struct{})
	limiter.buffer = make([]*executeJob, 0)

	return limiter
}

func (l *Limiter) AddJob(job *LimiterJob) {
	if l.started.CompareAndSwap(true, true) {
		return
	}
	if len(l.jobInChan) >= cap(l.jobInChan) {
		if l.strategy == Rejected {
			return
		} else if l.strategy == Accepted {
			l.buffer = append(l.buffer, &executeJob{
				wg:  &l.wg,
				job: job,
			})
		}
		return
	} else {
		l.jobInChan <- &executeJob{
			wg:  &l.wg,
			job: job,
		}
	}
}

func (l *Limiter) Start() []LimiterJobResult {
	l.started.Store(true)

	executeResult := make([]LimiterJobResult, 0)
loop:
	for {
		select {
		case execJob := <-l.jobInChan:
			{
				go func() {
					defer func() {
						if err := recover(); err != nil {
							executeResult = append(executeResult, LimiterJobResult{
								Error: fmt.Errorf("execute job panic, err = %s", err),
							})
						}
						l.mtx.Unlock()
						execJob.wg.Done()
						l.jobDownChan <- struct{}{} // 不管是否执行成功，都结束该任务
					}()

					execJob.wg.Add(1)
					l.mtx.Lock()
					result := execJob.job.Executor(execJob.job.Params)
					executeResult = append(executeResult, result)
				}()
			}
		case _ = <-l.jobDownChan:
			{
				if len(l.buffer) >= 1 {
					l.jobInChan <- l.buffer[0]
					l.buffer = l.buffer[1:]
				}
			}
		default:
			{
				if len(l.buffer) <= 0 {
					close(l.jobInChan)
				}
				if len(l.jobInChan) <= 0 {
					break loop
				}
			}
		}
	}
	l.wg.Wait()

	return executeResult
}
