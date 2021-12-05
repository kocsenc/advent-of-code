package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type command struct {
	name     string
	distance int
}

func readCommandsFromFile(path string) ([]command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commandSlice := strings.SplitN(scanner.Text(), " ", 2)

		commandDistance, err := strconv.Atoi(commandSlice[1])
		if err != nil {
			panic("Could not convert command distance to Integer")
		}

		command := command{commandSlice[0], commandDistance}
		commands = append(commands, command)
	}

	return commands, scanner.Err()
}

func main() {
	const inputFileName = "input.txt"
	commands, err := readCommandsFromFile(inputFileName)
	check(err)

	var horizontalPosition int
	var depth int
	var aim int

	for _, command := range commands {
		switch command.name {
		case "forward":
			horizontalPosition += command.distance
			depth += (aim * command.distance)
		case "up":
			aim -= command.distance
		case "down":
			aim += command.distance
		}
	}

	positionProduct := horizontalPosition * depth
	fmt.Printf("Horizontal Position: %d\nDepth Position: %d\nPosition Product: %d\n",
		horizontalPosition,
		depth,
		positionProduct)
}
