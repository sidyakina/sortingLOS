package main

import "fmt"

func main() {
	list := NewLOS(8)
	fmt.Printf("Old list: %v", *list)
	list.Sort1()
	fmt.Printf("\nAfter sort1: %v", *list)
	NewLOS(8).Sort2()
	fmt.Printf("\nAfter sort2: %v", *list)
}
