// https://adventofcode.com/2019/day/1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input file. %v.", err.Error())
	}
	defer func() {
		err = inputFile.Close()
		if err != nil {
			log.Fatalf("Unable to close input file. %v.", err.Error())
		}
	}()

	totalFuel := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		text := scanner.Text()
		mass, err := strconv.Atoi(text)
		if err != nil {
			log.Fatalf("Unable to convert mass value [%s] to integer. %v.", text, err.Error())
		}

		fuel := mass/3 - 2
		totalFuel += fuel

		for fuel > 0 {
			fuel = fuel/3 - 2

			if fuel > 0 {
				totalFuel += fuel
			}
		}
	}

	fmt.Println(totalFuel)
}
