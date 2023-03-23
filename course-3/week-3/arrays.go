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
	sort.Ints(slice)

	// write to channelOut
	channelOut <- slice

	// sort the array
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

func main() {

	data := []int{20, 30, 15, 50, 10, 60, 70, 40, 1, 50, 33, 90, -1}
	channelIn := make(chan []int, 4)
	channelOut := make(chan []int, 4)

	// create 4 workers
	workersCount := 4
	var workers []func(*sync.WaitGroup, chan []int, chan []int)
	for i := 0; i < workersCount; i++ {
		workers = append(workers, sortArray)
	}
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
	println("Sorted Array: ", toDisplayList(sortedData))

}
