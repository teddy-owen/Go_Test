package main

import "fmt"

func main() {
	var x int = 5;
	fmt.Println("hello, world!")
	fmt.Println(x)

	for i := 0; i < 50; i++ {
		fmt.Printf("%d\n",x)
		x = x * x 		
	}
}
