package main

import (
	"fmt"
)

func main() {
	fmt.Println("before sort array of values ")
	numbers := []int{2,3,4,6,1}
	fmt.Println(numbers)
	 
	SortOrderNumber(numbers)
}

func SortOrderNumber(numbers []int){
		fmt.Println("SortOrderNumber inside ",numbers)

	for i :=1;i<len(numbers);i++{
	   if numbers[i]<numbers[i-1]{
	   numbers[i] = numbers[i] + numbers[i-1]
	   numbers[i-1] =numbers[i]-numbers[i-1]
	   numbers[i] =numbers[i]-numbers[i-1]
	   i=0
	   }
	}
	fmt.Println("after sorting asc array of values ",numbers)
}
