package main

import (
	bt "bazel-tutorial/projects/main/greet"
	"fmt"
	"testing"
)

type BuildHelloTestInput struct {
	output string
}

func TestSayHello(t *testing.T) {

	t.Run("Testing Hello", func(t *testing.T) {

		output := ""

		for i, tt := range []BuildHelloTestInput{
			{
				output: "Hello Kenny!",
			},
		} {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				output = bt.SayHello()
				if tt.output != output {
					t.Errorf("Expected %v, but received %v", tt.output, output)
				}
			})
		}
	})
}
