package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
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
			filters := []filter{
				part1, byrFilter, hgtFilter, eyrFilter,
				hclFilter, eclFilter, pidFilter, iyrFilter,
			}
			if exec(passport, filters) {
				valid++
			}

		default:
		}
		if end {
			break
		}
	}
	fmt.Println("The number of valid passports: ", valid)
}

func exec(m map[string]string, filters []filter) bool {
	for _, f := range filters {
		if !f(m) {
			return false
		}
	}
	return true
}

type filter func(map[string]string) bool

func byrFilter(m map[string]string) bool {
	return durationFilter(m, "byr", 2002, 1920)
}

func iyrFilter(m map[string]string) bool {
	return durationFilter(m, "iyr", 2020, 2010)
}

func eyrFilter(m map[string]string) bool {
	return durationFilter(m, "eyr", 2030, 2020)
}

func hgtFilter(m map[string]string) bool {
	hgt, ok := m["hgt"]
	if !ok {
		return false
	}
	if len(hgt) <= 2 {
		return false
	}
	number, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		return false
	}
	switch hgt[len(hgt)-2:] {
	case "cm":
		return compare(number, 193, 150)
	case "in":
		return compare(number, 76, 59)
	default:
		return false
	}
}

func hclFilter(m map[string]string) bool {
	hcl, ok := m["hcl"]
	if !ok {
		return false
	}
	if !strings.HasPrefix(hcl, "#") {
		return false
	}
	if len(hcl) != 7 {
		return false
	}
	_, err := hex.DecodeString(hcl[1:])
	if err != nil {
		return false
	}
	return true
}

func eclFilter(m map[string]string) bool {
	ecl, ok := m["ecl"]
	if !ok {
		return false
	}
	for _, color := range []string{"amb", "blu", "brn",
		"gry", "grn", "hzl", "oth"} {

		if color == ecl {
			return true
		}
	}
	return false
}

func pidFilter(m map[string]string) bool {
	pid, ok := m["pid"]
	if !ok {
		return false
	}
	if len(pid) != 9 {
		return false
	}
	_, err := strconv.Atoi(pid)
	if err != nil {
		return false
	}
	return true
}

func durationFilter(m map[string]string, key string, max, min int) bool {
	s, ok := m[key]
	if !ok {
		return false
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return compare(v, max, min)
}

func compare(current, max, min int) bool {
	return !(current > max || current < min)
}

func part1(passport map[string]string) bool {
	if len(passport) == 8 {
		return true
	}
	if _, ok := passport["cid"]; !ok && len(passport) == 7 {
		return true
	}
	return false
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
