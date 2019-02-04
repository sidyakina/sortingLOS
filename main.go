package main

import "fmt"

func main() {
	list := NewLOS(8)
	fmt.Printf("Old list: %v", *NewLOS(8))
	list.Sort1()
	fmt.Printf("\nAfter sort1: %v", *list)
	list = NewLOS(8)
	list.Sort2()
	fmt.Printf("\nAfter sort2: %v", *list)
	list = NewLOS(8)
	list.Sort3()
	fmt.Printf("\nAfter sort3: %v", *list)
	list = NewLOS(8)
	list.Sort4()
	fmt.Printf("\nAfter sort4: %v", *list)
}
