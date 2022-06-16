package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func job(id int, w *workerPoll) {
	start := time.Now()
	fmt.Printf("task %d started \n", id)
	rand.Seed(start.UnixMilli())
	randnum := rand.Intn(100) + 100
	time.Sleep(time.Duration(randnum * int(time.Millisecond)))
	fmt.Println("task", id, "finished in", time.Since(start))
	w.complete <- 1
}

type workerPoll struct {
	totalTasks int
	ch         chan int
	complete   chan int
	res        int
	sigChan    chan os.Signal
}

func initWorkerPool(totalTasks int) workerPoll {
	ch := make(chan int, runtime.NumCPU())
	complete := make(chan int)
	sigChan := make(chan os.Signal)
	return workerPoll{totalTasks: totalTasks, ch: ch, complete: complete, sigChan: sigChan}
}

func (w *workerPoll) runPool() {
	stopChan := make(chan int)
	res := 0
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(context.Context) {
		signal.Notify(w.sigChan, os.Interrupt)
		for i := 1; i <= w.totalTasks; i++ {
			if ctx.Err() != nil {
				stopChan <- 1
				return
			}
			w.ch <- 1
			go job(i, w)
		}
	}(ctx)

	for {
		select {
		case <-w.sigChan:
			cancel()
			fmt.Println("shutting down the app")
		case v := <-w.complete:
			res += v
			if res == w.totalTasks {
				return
			}
			<-w.ch

		case <-stopChan:
			fmt.Println("total number of routines was completed", res)
			return
		}
	}

}

func main() {

	myPool := initWorkerPool(100)
	myPool.runPool()
	fmt.Println("EXIT")

}
