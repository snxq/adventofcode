package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := make(chan string)
	go readData(c)

	var end bool

	slopes := []*slope{{
		right: 1, down: 1,
	}, {
		right: 3, down: 1,
	}, {
		right: 5, down: 1,
	}, {
		right: 7, down: 1,
	}, {
		right: 1, down: 2,
	}}

	for {
		select {
		case line := <-c:
			if line == "" {
				end = true
				break
			}
			for _, slope := range slopes {
				slope.slip(line)
			}
		default:
			//
		}
		if end {
			close(c)
			break
		}
	}

	var result int = 1
	for _, slope := range slopes {
		// part1
		fmt.Printf("Right: %d, Down: %d, Square: %d, Tree: %d, X: %d, Y: %d\n",
			slope.right, slope.down, slope.square, slope.tree, slope.x, slope.y)
		// part2
		result = result * slope.tree
	}
	fmt.Printf("Part2 result: %d\n", result)
}

// different slopes
type slope struct {
	x, y         int
	right, down  int
	square, tree int
}

func (s *slope) slip(line string) {
	defer func() {
		s.x = s.x + s.right
		s.y++
	}()

	if s.x == 0 {
		return
	}
	if (s.y % s.down) != 0 {
		s.x = s.x - s.right
		return
	}
	if line[s.x%len(line)] == '#' {
		s.tree++
	} else {
		s.square++
	}
	return
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
