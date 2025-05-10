package greetings

import "fmt"

func SayHello() {
	fmt.Println("Hello World!")
}

func Greet(name string) string {
	return "Hello " + name + "!"
}
