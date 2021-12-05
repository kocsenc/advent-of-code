package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/asecurityteam/rolling"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readDepthMeasurementsFromFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var depths []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, e := strconv.Atoi(scanner.Text())
		if e != nil {
			panic("Could not Parse integer from depth measurement")
		}
		depths = append(depths, depth)
	}

	return depths, scanner.Err()
}

func main() {
	const inputFileName = "input.txt"
	const scanWindow = 3

	depths, err := readDepthMeasurementsFromFile(inputFileName)
	check(err)

	var rollingWindow = rolling.NewPointPolicy(rolling.NewWindow(scanWindow))

	for i := 0; i < scanWindow; i = i + 1 {
		rollingWindow.Append(float64(depths[i]))
	}

	previousRollingWindowSum := rollingWindow.Reduce(rolling.Sum)
	var increases int
	var decreases int

	for _, depth := range depths[scanWindow:] {
		rollingWindow.Append(float64(depth))

		peekRollingWindowSum := rollingWindow.Reduce(rolling.Sum)
		if peekRollingWindowSum > previousRollingWindowSum {
			increases++
		} else {
			decreases++
		}

		previousRollingWindowSum = peekRollingWindowSum
	}

	fmt.Println(increases)
	fmt.Println(decreases)
}
