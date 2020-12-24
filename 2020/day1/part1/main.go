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

	result := 1
	for i, x := range data {
		y := 2020 - x
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

	fmt.Println(result)
}

func readData() ([]int, error) {
	b, err := ioutil.ReadFile("2020/day1/part1/data")
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
