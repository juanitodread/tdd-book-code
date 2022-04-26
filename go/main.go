package main

import (
	"fmt"
	"os"
	"tdd/hello"
	"tdd/samples"
)

func misc() {
	var name = "Juan"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	message := hello.Hello(name)
	fmt.Println(message)

	res := hello.FizzBuzz("hello")
	if res != nil {
		fmt.Println(*res)
	} else {
		fmt.Println("res is nil")
	}

	res2 := hello.FizzBuzzRef(nil)
	if res2 != nil {
		fmt.Println(*res2)
	} else {
		fmt.Println("res is nil")
	}

	myMap := map[string]float64{
		"one": 1.0,
		"two": 2.0,
	}

	fmt.Println(myMap)

	var myPointer *map[string]float64 = &myMap

	fmt.Println(myPointer)
	fmt.Printf("This is *myPointer: %+v\n", *myPointer)

	for key, val := range myMap {
		fmt.Printf("[%s]=%f\n", key, val)
	}

	for index, char := range "a slice of chars" {
		fmt.Printf("[%d:%c],", index, char)
	}
}

func main() {
	// misc()
	samples.Structs()
	//samples.Concurrency()
}
