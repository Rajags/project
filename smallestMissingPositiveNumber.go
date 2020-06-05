package main

import (
	"fmt"
)
// find smallest mising positive number 
func smallestMissingPositiveNumber(array []int)int{
	smallestval := make(map[int]int)
	positivenumber :=1
	present := false
	for _,val := range array{
		smallestval[val] = val
	}
	
	fmt.Println("smallestval",smallestval)
	
	for !present{
		if _,ok:=smallestval[positivenumber]; !ok{		
		    present = true
		}else {
			positivenumber++
		}
	}

return positivenumber
}

func main() {
	a := []int{1, 4, 2, -1, 0, 2, 5 }
	fmt.Println(smallestMissingPositiveNumber(a))
	
}
