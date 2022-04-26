package samples

import (
	"fmt"
	"time"
)

type Results struct {
	Users     string
	Addresses string
	Cars      string
}

type LongRunningTask func() string

func GetUsers() string {
	return "users-data"
}

func GetAddresses() string {
	return "addresses-data"
}

func GetCars() string {
	return "cars-data"
}

func run(seconds int, task string, fn LongRunningTask) string {
	fmt.Printf("Calling function: %+v\n", task)

	time.Sleep(time.Duration(seconds) * time.Second)

	fmt.Printf("Finished function: %+v\n", task)
	return fn()
}

func SimpleChannels() {
	ch := make(chan string) // create string channel

	go func() { // go routine using anonymous function
		message := <-ch // Blocking receive; assigns to message
		fmt.Println(message)
		ch <- "pong"
	}()

	ch <- "ping"
	fmt.Println(<-ch)
}

func LRTsSecuential() {
	users := run(3, "GetUsers", GetUsers)
	addresses := run(2, "GetAddresses", GetAddresses)
	cars := run(4, "GetCars", GetCars)

	fmt.Println("LRTs finished", users, addresses, cars)
}

func LRTs() {
	users := make(chan string)
	addresses := make(chan string)
	cars := make(chan string)

	go func() {
		users <- run(2, "GetUsers", GetUsers)
	}()

	go func() {
		addresses <- run(3, "GetAddresses", GetAddresses)
	}()

	go func() {
		cars <- run(5, "GetCars", GetCars)
	}()

	fmt.Println("LRTs finished", users, addresses, cars) // Executed immediatly

	// let's process results when goroutines finished
	results := new(Results)
	results.Users = <-users
	results.Addresses = <-addresses
	results.Cars = <-cars

	// Another way to wait for results
	// for i := 0; i < 3; i++ {
	// 	select {
	// 	case u := <-users:
	// 		results.Users = u
	// 	case a := <-addresses:
	// 		results.Addresses = a
	// 	case c := <-cars:
	// 		results.Cars = c
	// 	}
	// }

	fmt.Printf("Results: %+v\n", *results)
}

func Concurrency() {
	start := time.Now()
	//SimpleChannels()
	//LRTsSecuential()
	LRTs()
	fmt.Printf("Execution time: %d\n", time.Since(start)/time.Second)
}
