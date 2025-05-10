package main

import (
	"GO-BUILD-WITH-BAZEL/com/saif/hello-world/greetings"
	"fmt"
)

func main() {
	greetings.SayHello()

	fmt.Println(greetings.Greet("M. Geremy"))
}
