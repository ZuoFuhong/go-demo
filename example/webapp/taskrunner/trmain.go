package taskrunner

import (
	"time"
)

/*
	scheduler调度器，每3s执行一次
*/

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	r := NewRunner(30, true, ClearDispatcher, ClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}
