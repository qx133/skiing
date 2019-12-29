package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadFile parse the input of a file containing arrays of int
func ReadFile(path string) ([]int, [][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("unable to open input file")
	}

	scanner := bufio.NewScanner(file)
	header := []int{}
	isHeader := true
	input := [][]int{}

	for scanner.Scan() {
		if isHeader {
			header = numbers(scanner.Text())
			isHeader = false
		} else {
			temp := []int{}
			temp = numbers(scanner.Text())
			input = append(input, temp)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("unable to read input file")
	}

	return header, input
}

func numbers(s string) []int {
	var res []int
	for _, v := range strings.Fields(s) {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("unable to parse int")
		}
		res = append(res, i)
	}
	return res
}
