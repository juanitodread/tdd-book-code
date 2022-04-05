package hello

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// References are declared with * => *string is a string reference
// as reference you can verify if is null => nil
// you can get its value with *variable => *word
// you can wrap a value as reference with & => &word

func FizzBuzz(word string) *string {
	if word == "fizzbuzz" {
		return &word
	}

	return nil
}

func FizzBuzzRef(word *string) *string {
	if word != nil && *word == "fizzbuzz" {
		return word
	}

	return nil
}
