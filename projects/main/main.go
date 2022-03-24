package main

import (
	bt "bazel-tutorial/projects/main/greet"
	"fmt"
)

func main() {
	fmt.Println(bt.SayHello())
	fmt.Println(bt.Greet())
}
