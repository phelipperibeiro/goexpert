package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // slice from 10 position

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) // slicing the slice from index 0 to 0

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) // slicing the slice from index 0 to 4

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // slicing the slice from index 2 to the end

	s = append(s, 10) // append 10 position to the slice

	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2]) // slicing the slice from index 0 to 2

}

// go run 3-foundation/6-slices/*.go
