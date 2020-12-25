package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	var sum int = 2020

	fmt.Println("part1 result: ", part1(data, sum))
	fmt.Println("part2 result: ", part2(data, sum))
}

// part1 x+y=2020 x*y=?
func part1(data []int, sum int) int {
	result := 1
	for i, x := range data {
		y := sum - x
		var exist bool
		for _, z := range data[i:] {
			if z == y {
				exist = true
				break
			}
		}
		if exist {
			result = result * (x * y)
		}
	}

	return result
}

// part2 x+y+z=2020 x*y*z=?
func part2(data []int, sum int) int {
	result := 1
	for _, x := range data {
		y := sum - x
		p := part1(data, y)
		if p != 1 {
			result = result * x * p
			break
		}
	}

	return result
}

func readData() ([]int, error) {
	b, err := ioutil.ReadFile("2020/day1/data")
	if err != nil {
		return nil, err
	}
	var data []int
	for _, item := range bytes.Split(b, []byte("\n")) {
		d, err := strconv.Atoi(string(item))
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}
