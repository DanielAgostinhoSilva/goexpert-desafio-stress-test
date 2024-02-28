package infrastructure

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Status int
	Err    error
}

type StressTestReport struct {
	totalExecutionTime    time.Duration
	maxNumberRequest      int
	maxNumberRequestError int
	statuses              map[int]int
}

func NewStressTestReport() *StressTestReport {
	return &StressTestReport{statuses: make(map[int]int)}
}

func (s *StressTestReport) Execute(url string, totalRequest, concurrentRequest int) {
	results := make(chan Result, totalRequest)
	done := make(chan bool)
	var wg sync.WaitGroup

	start := time.Now()

	fetchPerWorker := totalRequest / concurrentRequest
	for i := 0; i < concurrentRequest; i++ {
		wg.Add(1)
		go s.worker(url, fetchPerWorker, results, &wg, done)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	progressBar := 0
	for {
		select {
		case r := <-results:
			if r.Err != nil {
				s.maxNumberRequestError++
			} else {
				s.maxNumberRequest++
				s.statuses[r.Status]++
			}
		case <-done:
			progressBar++
			s.displayProgressBar(progressBar, totalRequest)
			if progressBar == totalRequest {
				fmt.Println()
			}
		}
		if progressBar == totalRequest {
			break
		}
	}

	s.totalExecutionTime = time.Since(start)
	s.printReport()
}

func (s *StressTestReport) worker(url string, n int, r chan<- Result, wg *sync.WaitGroup, done chan<- bool) {
	defer wg.Done()

	client := http.Client{}
	for i := 0; i < n; i++ {
		resp, err := client.Get(url)
		if err != nil {
			r <- Result{Err: err}
		} else {
			resp.Body.Close()
			r <- Result{Status: resp.StatusCode}
		}
		done <- true
	}
}

func (s *StressTestReport) printReport() {
	fmt.Printf("Tempo total gasto na execução: %v\n", s.totalExecutionTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", s.maxNumberRequest)
	fmt.Printf("Quantidade total de requests com erro: %d\n", s.maxNumberRequestError)
	for status, count := range s.statuses {
		fmt.Printf("Quantidade total de requests com status %d: %d\n", status, count)
	}
}

func (s *StressTestReport) displayProgressBar(completed, totalRequest int) {
	fmt.Printf("\r[%-50s] %d%%", string(repeat('#', completed*50/totalRequest)), completed*100/totalRequest)
}

func repeat(char rune, count int) []rune {
	var result []rune
	for i := 0; i < count; i++ {
		result = append(result, char)
	}
	return result
}
