package user

import (
	"sync"
)

type Job interface {
	DispatchWorkersCreateUserBulk_(jobs <-chan []interface{}, wg *sync.WaitGroup)
}

type job struct {
	service Service
}

func NewJob(service Service) *job {
	return &job{service}
}

const totalWorker = 20

func (j *job) DispatchWorkersCreateUserBulk_(jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				j.service.CreateUserBulk(workerIndex, counter, job)
				wg.Done()
				counter++
			}
		}(workerIndex, jobs, wg)
	}
}
