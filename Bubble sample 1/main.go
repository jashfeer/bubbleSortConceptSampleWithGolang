package main

import "fmt"

func main() {
	array := []int{20, 3, 5,-2,-30,-100, 70, 10, 55}
	arr := bubbleSort(array)
	fmt.Println(arr)

}

func bubbleSort(array []int) []int {

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
