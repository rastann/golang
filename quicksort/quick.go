package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_arr(size int) []int {
	arr := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(999) - rand.Intn(999)
	}
	return arr
}

func quicksort(arr []int, start, end int) {
	if start < end {
		p := partition(arr, start, end)
		quicksort(arr, start, p-1)
		quicksort(arr, p+1, end)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	//slice := generate_arr(10)
	slice := []int{12,324,32,1,232,4,10}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice, 0, len(slice)-1)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
