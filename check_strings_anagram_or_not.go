package main

import (
	"fmt"
)

//Implement the method to check the given strings are anagram. 
//An anagram of a string is another string that contains same characters, 
//only the order of characters can be different. 
//Sample:
//Input: abcc,cbca Output: true
//Input: abcc,cab  Output: false

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}
	for _, val := range s1{
		if !containString(string(val),s2) {
			return false
		}
	}
	
	return true
}

func containString(element, str string)(bool){
	for _,v := range str{		
		if element == string(v){
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Check if Two strings are anagram")
	fmt.Println("1.", IsAnagram("listen", "silent"))
	fmt.Println("2.", IsAnagram("emerio", "mario"))
}
