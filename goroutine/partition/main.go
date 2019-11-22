package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type (
    app struct {
        threads     int
        workerMutex sync.Mutex
        workers     []*worker
    }

    worker struct {
        in  chan job
        out chan error
    }

    job func() error
)

func main() {
    app := app{
        threads:     5,
        workerMutex: sync.Mutex{},
        workers:     make([]*worker, 5),
    }

    job := func(msg string) job {
        return func() error {
            fmt.Println("job running:", msg)
            time.Sleep(1 * time.Second)
            return nil
        }
    }

    ctx := context.Background()

    for i := 0; i < 10; i++ {
        worker := app.getWorker(ctx, i)
        go func(i int) {
            worker.in <- job(fmt.Sprintf("%v", i))
            <-worker.out
        }(i)
    }

    time.Sleep(2 * time.Second)
}

func (re *app) getWorker(ctx context.Context, id int) *worker {
    re.workerMutex.Lock()
    defer re.workerMutex.Unlock()

    workerId := id % re.threads

    if nil == re.workers[workerId] {
        re.workers[workerId] = &worker{
            in:  make(chan job, 1),
            out: make(chan error, 1),
        }

        go func() {
            for {
                select {
                case job := <-re.workers[workerId].in:
                    err := job()
                    re.workers[workerId].out <- err

                case <-ctx.Done():
                    re.workers[workerId].out <- ctx.Err()
                    break
                }

            }
        }()
    }

    return re.workers[workerId]
}
