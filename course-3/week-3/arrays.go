package main

import (
	"fmt"
	"sort"
	"sync"
)

// / The program should partition the array into 4 parts, each of which is sorted by a different goroutine
// Each partition should be of approximately equal size.
//
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

func sortArray(group *sync.WaitGroup, channelIn, channelOut chan []int) {

	// read from channelIn
	slice := <-channelIn
	id := fmt.Sprintf("%p", slice)
	fmt.Printf("Worker[%s]: I'll work with <%s>\n", id, toDisplayList(slice))

	// sort the array
	sort.Ints(slice)
	fmt.Printf("Worker[%s]: I'm done with <%s>\n", id, toDisplayList(slice))

	// write to channelOut
	channelOut <- slice
	group.Done()
}

func toDisplayList(data []int) string {

	str := "["

	for _, v := range data {
		str += fmt.Sprintf(" %d", v)
	}

	str += " ]"

	return str
}

// requestData Prompts the user to enter a number of positive integers
func requestData(size int) []int {

	var data []int
	var input int

	fmt.Printf("Main: Please enter %d positive integers.\n", size)
	for i := 0; i < size; i++ {
		fmt.Printf("[%d]> ", i+1)
		_, err := fmt.Scan(&input)

		if err != nil {
			println("Error: ", err)
			continue
		}

		data = append(data, input)
	}

	return data
}

// createWorkers creates a slice of workers
func createWorkers(size int) []func(*sync.WaitGroup, chan []int, chan []int) {
	var workers []func(*sync.WaitGroup, chan []int, chan []int)
	for i := 0; i < size; i++ {
		workers = append(workers, sortArray)
	}
	return workers
}

func main() {

	// Given
	data := requestData(12)
	channelIn := make(chan []int, 4)
	channelOut := make(chan []int, 4)

	// create 4 workers
	workersCount := 4
	workers := createWorkers(workersCount)

	var wg sync.WaitGroup
	for _, worker := range workers {
		wg.Add(1)
		go worker(&wg, channelIn, channelOut)
	}

	// Send data to channelIn
	size := len(data) / workersCount
	for i := 0; i < workersCount; i++ {

		start := i * size
		end := (i + 1) * size

		// Adjust the last partition
		if end > len(data) || i == workersCount-1 {
			end = len(data)
		}

		channelIn <- data[start:end]
	}

	// Receive data from channelOut
	var sortedData []int
	wg.Wait()
	for i := 0; i < workersCount; i++ {
		sortedData = append(sortedData, <-channelOut...)
	}

	// sort the array
	sort.Ints(sortedData)
	println("Main: result=", toDisplayList(sortedData))
}
