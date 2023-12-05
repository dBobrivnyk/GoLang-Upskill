package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func StartGoroutine(i int) {
	time.Sleep(20 * time.Millisecond)
	fmt.Println("goroutine:", i)
}

// channel is already safe
type SafeChannel struct {
	mu sync.Mutex
	v  chan int
}

func (safeCh *SafeChannel) Write(i int, wg *sync.WaitGroup) {
	//no need to use lock, reading and writing from channels is safe
	safeCh.mu.Lock()
	safeCh.v <- i
	safeCh.mu.Unlock()
}

func (safeCh *SafeChannel) Read(i int, wg *sync.WaitGroup) {
	for {
		v, ok := <-safeCh.v
		if ok == true {
			fmt.Println(i, "consumer reading:", v)
		} else {
			break
		}
	}
	wg.Done()
}

func ProducerConsumer() {
	wg := sync.WaitGroup{}
	ch := SafeChannel{
		mu: sync.Mutex{},
		v:  make(chan int),
	}

	wg.Add(1)
	go startProducer(&ch, &wg)

	for i := 1; i <= 15; i++ {
		wg.Add(1)
		go ch.Read(i, &wg)
	}

	wg.Wait()
	fmt.Println("...Finish...")

}

func startProducer(ch *SafeChannel, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		fmt.Println("Writing", i, "to channel")
		ch.Write(i, wg)
	}
	fmt.Println("...Producer`s work finished...")
	close(ch.v)
}
