package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("pin pan")
		case i%3 == 0:
			fmt.Println("pin")
		case i%5 == 0:
			fmt.Println("pan")
		default:
			fmt.Println(i)
		}
	}
}
