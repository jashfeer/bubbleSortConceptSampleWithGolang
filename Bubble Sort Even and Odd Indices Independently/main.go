package main

import "fmt"
func main(){

	array:=[]int{10,11,12,13,2,21,45,32,45,98,43,20,31,36}
	arr:=sortEvenOdd(array)
	fmt.Println(arr)



}
func sortEvenOdd(array []int) []int {
    for i:=0;i<len(array)-1;i++{
        for j:=0;j<len(array)-2;j++{
            if j%2==0{
                if array[j]>array[j+2]{
                    array[j],array[j+2]=array[j+2],array[j]
                }
            }else{
               if array[j]<array[j+2]{
                    array[j],array[j+2]=array[j+2],array[j]
            }
        }
     }
    }
        return array
}