package main

import (
	"fmt"

	"github.com/qx133/main/internal/input"
	"github.com/qx133/main/pkg/skiing"
)

func main() {
	inputs := []string{
		"./sample_input/input1.txt",
		"./sample_input/input2.txt",
	}

	fmt.Println("length", "height")
	for _, v := range inputs {
		_, inputFile := input.ReadFile(v)
		length, height := skiing.BestSkiPath(inputFile)
		fmt.Println(length, height)
	}
}
