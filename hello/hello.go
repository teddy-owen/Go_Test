package main

import (
	"fmt"
	"errors"
	"github.com/teddy-owen/Go_Test/stringutil"
)

type node struct {
	value int
	next *node
}

func main() {
	var x int = 5;
	fmt.Println("hello, world!")
	fmt.Println(x)

	for i := 0; i < 50; i++ {
		fmt.Printf("%d\n",x)
		x = x * x 		
	}

	
	fmt.Println(stringutil.Reverse("Hello, Go!"))
	fmt.Println(stringutil.Underscore("Hello, Go!"))
	stringutil.Ascii("Hello, Go!")	
	
	var n1 node
	var n2 node

	n1.value = 1
	n2.value = 2

	n1.next = &n2
	
	fmt.Println(n1)
	fmt.Println(n2)
	
	v1,e1 := getNextValue(n1)
	v2,e2 := getNextValue(n2)

	if e1 == nil {
		fmt.Printf("n1.next -> %d\n",v1)
	}else{
		fmt.Println(e1)
	}
	if e2 == nil {
		fmt.Printf("n2.next -> %d\n",v2)
	}else{
		fmt.Println(e2)
	}

}

func getNextValue(n node)(value int,err error){
	if n.next != nil {
		return (*n.next).value, nil
	}else{
		return -1, errors.New("No next node")
	}
}


