package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

type Numbers struct {
	A int `json:"a"`
	B int `json:"b"`
}

func sum(numbers []Numbers, start, end int, ch chan int) {
	sum := 0
	for i := start; i < end; i++ {
		sum += numbers[i].A + numbers[i].B
	}
	ch <- sum
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run sum_calculator.go <input_file> <num_goroutines>")
		return
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage: go run sum_calculator.go <input.json> <num_goroutines>")
		return
	}

	if os.Args[2] == "0" {
		fmt.Println("Error: num_goroutines cannot be 0")
		return
	}

	numGoroutines, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	inputFile := os.Args[1]

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("Error: %v file does not exist\n", inputFile)
		return
	}

	file, err := ioutil.ReadFile(inputFile)
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
		go func(start, end int) {
			defer wg.Done()
			sum(numbers, start, end, ch)
		}(start, end)
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
