package main

import (
	"fmt"
	"time"
)

type Jober interface {
	Do()
}

type Worker struct {
	JobChan chan Jober
}

type WorkerPool struct {
	Workerlen int
	JobChan chan Jober
	WorkerChan chan chan Jober
}

func NewWorker() *Worker {
	return &Worker{make(chan Jober)}
}

func NewWorkerPool(max int) *WorkerPool {
	return &WorkerPool{
		Workerlen: max,
		JobChan: make(chan Jober),
		WorkerChan: make(chan chan Jober, max),
	}
}

func (this *Worker) Run(wc chan chan Jober) {
	go func() {
		fmt.Println("worker ... ")
		for {
			wc <- this.JobChan
			select {
			case job := <- this.JobChan:
				job.Do()
			}
		}
	}()
}

func (this *WorkerPool) Run() {
	fmt.Println("init worker ...")
	for i := 0; i < this.Workerlen; i++ {
		w := NewWorker()
		w.Run(this.WorkerChan)
	}
	go func() {
		for {
			select {
			case job := <- this.JobChan:
				w := <- this.WorkerChan
				w <- job
			}
		}
	}()
}

func (this *Score) Do() {
	fmt.Println("do ... ")
}

type Score struct {
	Num int
}

func main() {
	num := 5
	wp := NewWorkerPool(num)
	wp.Run()
	fmt.Println(wp)
	for i := 0; i< 5; i++ {
		sc := new(Score)
		wp.JobChan <- sc
	}
	for {
		time.Sleep(2 * time.Second)
	}
}