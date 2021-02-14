package main 

import (
	"fmt"
	
)

func main() {
	fmt.Print("Enter number of elements: ")
	var num int
	fmt.Scan(&num)
	
	arr := []int{}

	for i := 0; i < num; i++ {
		fmt.Print("Element ",  i+1, ": ")
		var tmp int
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}

	quickSort(arr, 0, len(arr) - 1)
	printArr(arr)
}

func quickSort(arr []int, low int, high int) {
	if (low < high) {

		pi := partition(arr, low, high)
		
		quickSort(arr, low, pi - 1)
		quickSort(arr, pi + 1, high) 

	}
}

func partition(arr []int, low int, high int) int {

	i := low
	pivot := arr[high]

	for j := low; j < high; j++ {
		if (arr[j] < pivot) {
			arr[j], arr[i] = swap(arr[j], arr[i])
			i++
		}
	}
	arr[high], arr[i] = swap(arr[high], arr[i])

	return i
} 

func swap(a int, b int) (int, int) {
	return b, a
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
	fmt.Print("\n")
}