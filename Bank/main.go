package main

import "fmt"

func main() {
	
	categories := []string{"auto", "food", "beauty", "mobile", "fun"}
	top3:=categories[0:3]
	fmt.Println(top3)
	fmt.Println(len(top3))
	fmt.Println(cap(top3))
	
	/*
	categories := []string{"auto", "food", "beauty", "mobile", "fun"}
	top3:=categories[0:3]
	fmt.Println(top3)
	categories[1]:="it"
	fmt.Println(cap(top3))
	*/

}
