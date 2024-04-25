package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

const (
	fileName = "input.json"
)

type Numbers struct {
	A int `json:"a"`
	B int `json:"b"`
}

func sum(numbers []Numbers, start, end int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	sum := 0
	for i := start; i < end; i++ {
		sum += numbers[i].A + numbers[i].B
	}
	ch <- sum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run sum_calculator.go <num_goroutines>")
		return
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage: go run sum_calculator.go <num_goroutines>")
		return
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("Error: input.json file does not exist")
		return
	}

	if os.Args[1] == "0" {
		fmt.Println("Error: num_goroutines cannot be 0")
		return
	}

	numGoroutines, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var numbers []Numbers
	err = json.Unmarshal(file, &numbers)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	totalNumbers := len(numbers)
	numPerRoutine := totalNumbers / numGoroutines

	var wg sync.WaitGroup
	ch := make(chan int, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		start := i * numPerRoutine
		end := start + numPerRoutine
		if i == numGoroutines-1 {
			end = totalNumbers
		}
		go sum(numbers, start, end, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalSum := 0
	for sum := range ch {
		totalSum += sum
	}

	fmt.Println("Total sum:", totalSum)
}
