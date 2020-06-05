package main

import (
	"fmt"
)

func main() {
	fmt.Println("before reverse string : ")
	str := "test"
	reverseString(str)
}

func reverseString(str string){
	revstr := ""
	for _,val := range str{
		revstr = string(val)+revstr
	}
	fmt.Println("after reverse string : ",revstr)
  if str == revstr {
    fmt.Println("Is palindrome ",str)
  }else {
    fmt.Println("Not a palindrome",str)
  }
}
