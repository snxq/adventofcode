package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := make(chan string)
	go readData(c)

	var (
		end          bool
		x            int
		square, tree int
	)
	for {
		select {
		case line := <-c:
			if line == "" {
				end = true
				break
			}
			if x == 0 {
				x = x + 3
				continue
			}
			if line[x%len(line)] == '#' {
				tree++
			} else {
				square++
			}
			x = x + 3
		default:
			//
		}
		if end {
			break
		}
	}
	fmt.Printf("Square: %d, Tree: %d\n", square, tree)
}

func readData(c chan string) {
	file, err := os.Open("2020/day3/data")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	c <- ""
}
