package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	row1 := []int{}
	row2 := []int{}
	result := 0
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		row := strings.Split(line, "   ")
		n1, _ := strconv.Atoi(row[0])
		n2, _ := strconv.Atoi(row[1])
		row1 = append(row1, n1)
		row2 = append(row2, n2)
	}
	sort.Ints(row1)
	sort.Ints(row2)
	for i, v := range row1 {
		value := row2[i] - v
		value_abs := math.Abs(float64(value))
		result += int(value_abs)
	}
	fmt.Println(result)
}

func part2() {
	row1 := []int{}
	row2 := []int{}
	result := 0
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		row := strings.Split(line, "   ")
		n1, _ := strconv.Atoi(row[0])
		n2, _ := strconv.Atoi(row[1])
		row1 = append(row1, n1)
		row2 = append(row2, n2)
	}

	for _, v := range row1 {
		app := 0
		for _, v2 := range row2 {
			if v == v2 {
				app += 1
			}
		}
		result += v * app
	}
	fmt.Println(result)
}

func main() {
	part1()
	part2()
}
