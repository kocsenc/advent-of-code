package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var increases int
	var decreases int
	var previous int

	// Set Initial value
	scanner.Scan()
	previous, err = strconv.Atoi(scanner.Text())
	check(err)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		check(err)

		if num > previous {
			increases++
		} else {
			decreases++
		}

		previous = num
	}

	fmt.Println(increases)
	fmt.Println(decreases)
}
