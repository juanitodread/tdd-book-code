package samples

import "fmt"

type Friendly interface {
	SayHello() string
}

type Healthy interface {
	DoExcercise() string
}

// Interface embedding
type Smarter interface {
	Friendly
	Healthy
}

type Address struct {
	Street string
	Number int
}

// struct embedding
type User struct {
	Name    string
	Address Address
}

type User2 struct {
	Name string
	Address
}

type SmarterHuman struct {
	Name string
	Friendly
	Healthy
}

func (s SmarterHuman) SayHello() string {
	return "Hello, I'm friendly"
}

func (s SmarterHuman) DoExcercise() string {
	return "I'm doing excercise"
}

func (s *SmarterHuman) RefFunction() string {
	return "This is a function of SmarterHuman ref"
}

func Structs() {
	user := User{
		Name: "user-1",
		Address: Address{
			Street: "street-1",
			Number: 22,
		},
	}

	fmt.Printf("User: %+v\n", user)

	user2 := User2{
		Name: "user-2",
		Address: Address{
			Street: "street-2",
			Number: 24,
		},
	}

	fmt.Printf("User2: %+v\n", user2)

	var smarter Smarter
	smarter = SmarterHuman{
		Name: "smarter",
	}

	fmt.Printf("smarter: %+v, hello: %s, excercise: %s\n", smarter, smarter.SayHello(), smarter.DoExcercise())

	val, ok := smarter.(Friendly)
	fmt.Printf("Is smarter friendly?: %T, ok=%t\n", val, ok)

	smarterRef := &smarter
	fmt.Println("smarterRef", smarterRef)

	// Error, we need to cast to SmarterHuman in order to call the function RefFunction()
	// smarter.RefFunction()
	smarterHuman := smarter.(SmarterHuman)
	smarterHumanRef := &smarterHuman
	fmt.Println("smarterHumanRef.RefFunction()", smarterHumanRef.RefFunction())
	fmt.Println("smarterHumanRef.DoExcercise()", smarterHumanRef.DoExcercise()) // automatically compilers gets the value (*smarterHumanRef).DoExcercise()
	fmt.Println("smarterHuman.RefFunction()", smarterHuman.RefFunction())
}
