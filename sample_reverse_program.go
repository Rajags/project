package main

import (
	"fmt"
)

func main() {
	fmt.Println("reverse the string given string")
	str := "test"
	reverseString(str)
}

func reverseString(str string){
	var reversestr string
	for i := len(str)-1;i>=0;i--{
	   reversestr = reversestr + string(str[i])
	}
	fmt.Println(reversestr)
	revstr := ""
	for _,val := range reversestr{
		revstr = string(val)+revstr
	}
	fmt.Println("after reverse string ",revstr)
}
