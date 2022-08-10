package main

import "fmt"

func main() {
	arr := []int{20, 21, 30, 21, 50, 43, 44, 70,-20,-21, 72, 75}
	a := sortArrayByParity(arr)
	fmt.Print(a)

}

func sortArrayByParity(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		flag := 0
		for j := 0; j < len(array)-1; j++ {
			if array[j]%2 != 0 && array[j+1]%2 == 0  {
				array[j], array[j+1] = array[j+1], array[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
	return array

}
