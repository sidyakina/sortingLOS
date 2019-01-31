package main

import "fmt"

func main() {
	list := NewList(8)
	fmt.Printf("Old list: %v", *list)
	list.Sort()
	fmt.Printf("\nAfter sort: %v", *list)
}