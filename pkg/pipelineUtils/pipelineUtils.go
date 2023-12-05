package pipelineUtils

import (
	"fmt"
)

func NumberProducer(pipe chan<- int, dataToProduce []int) {
	for i := range dataToProduce {
		pipe <- dataToProduce[i]
	}
	close(pipe)
}

func TextProducer(pipe chan<- string, dataToProduce []string) {
	for i := range dataToProduce {
		pipe <- dataToProduce[i]
	}
	close(pipe)
}

// self study: https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308
func Merge2(textPipe <-chan string, numberPipe <-chan int) <-chan interface{} {
	var counter, tmp = cap(textPipe) + cap(numberPipe), 0
	res := make(chan interface{})

	go func() {
		for {
			select {
			case text, ok := <-textPipe:
				if ok == true {
					fmt.Println(text)
					res <- text
					tmp++
				}
			case number, ok := <-numberPipe:
				if ok == true {
					fmt.Println(number)
					res <- number
					tmp++
				}
			}
			if counter == tmp {
				break
			}
		}

		//should be deferd
		close(res)
	}()

	return res
}

func Merge(textPipe <-chan string, numberPipe <-chan int) <-chan interface{} {
	res := make(chan interface{})

	// both loops should be run in separate goroutines
	// to read in parallel
	// currently second loop will start after textPipe will be closed
	go func() {
		for n := range textPipe {
			res <- n
		}
		for n := range numberPipe {
			res <- n
		}
		//should be deferd
		close(res)
	}()

	return res
}

func StartMerging() {
	//wg := sync.WaitGroup{}
	textData := []string{"A", "B", "C"}
	numberData := []int{1, 2, 3}

	textChannel := make(chan string, 3)
	numberChannel := make(chan int, 3)

	//dump from text
	go TextProducer(textChannel, textData)

	//dump from int
	go NumberProducer(numberChannel, numberData)
	res := Merge2(textChannel, numberChannel)
	for n := range res {
		fmt.Println("Data showed in main:", n)
	}

}
