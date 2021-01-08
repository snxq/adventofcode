package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	c := make(chan *line)
	var (
		result1, result2 int
		end              bool
	)
	go readData(c)
	for {
		select {
		case l := <-c:
			if l == nil {
				end = true
				break
			}
			// part 1
			var times int
			for _, c := range l.password {
				if string(c) == l.key {
					times++
				}
			}
			if times >= l.minTimes && times <= l.maxTimes {
				result1++
			}
			// part 2
			if !((string(l.password[l.maxTimes-1]) == l.key) ==
				(string(l.password[l.minTimes-1]) == l.key)) {

				result2++
			}
		default:
			//
		}
		if end {
			break
		}
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

// Password every line
type line struct {
	maxTimes, minTimes int
	key                string
	password           string
}

func readData(c chan *line) {
	file, err := os.Open("2020/day2/data")
	checkErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := &line{}
		err := l.Unmarshal(scanner.Text())
		checkErr(err)
		c <- l
	}
	c <- nil
}

func (l *line) Unmarshal(s string) (err error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return errors.New("error line, " + s)
	}
	l.password = strings.TrimSpace(parts[1])
	timesKey := strings.Split(parts[0], " ")
	if len(timesKey) != 2 {
		return errors.New("error line, " + s)
	}
	l.key = timesKey[1]
	times := strings.Split(timesKey[0], "-")
	if len(times) != 2 {
		return errors.New("error line, " + s)
	}
	l.minTimes, err = strconv.Atoi(times[0])
	if err != nil {
		return err
	}
	l.maxTimes, err = strconv.Atoi(times[1])
	return err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
