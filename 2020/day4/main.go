package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	c := make(chan map[string]string)
	go readData(c)

	var (
		valid int
		end   bool
	)
	for {
		select {
		case passport := <-c:
			if passport == nil {
				end = true
				break
			}
			if len(passport) == 8 {
				valid++
				continue
			}
			if _, ok := passport["cid"]; !ok && len(passport) == 7 {
				valid++
			}

		default:
			//
		}
		if end {
			break
		}
	}
	fmt.Println("The number of valid passports: ", valid)
}

func readData(c chan map[string]string) {
	file, err := os.Open("2020/day4/data")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	passport := map[string]string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			c <- passport
			passport = map[string]string{}
			continue
		}
		columns := strings.Split(scanner.Text(), " ")
		for _, c := range columns {
			kv := strings.Split(c, ":")
			if len(kv) != 2 {
				panic("invalid kv")
			}
			passport[kv[0]] = kv[1]
		}
	}
	c <- passport
	c <- nil
}
