package main

import (
	"fmt"
	"os"
	"tdd/hello"
)

func main() {
	var name = "Juan"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	message := hello.Hello(name)
	fmt.Println(message)

	res := hello.FizzBuzz("hello")
	if res != nil {
		println(*res)
	} else {
		println("res is nil")
	}

	res2 := hello.FizzBuzzRef(nil)
	if res2 != nil {
		println(*res2)
	} else {
		println("res is nil")
	}
}
