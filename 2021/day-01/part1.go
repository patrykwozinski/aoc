package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	measurements := 0
	lastNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentNumber, _ := strconv.Atoi(scanner.Text())

		if lastNumber != 0 && currentNumber > lastNumber {
			measurements += 1
		}

		lastNumber = currentNumber
	}

	fmt.Println(measurements)
}
