package adventoftasks

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func countValidMulSubstrings(s string, instruction chan string) int {
	re := regexp.MustCompile(`(?:mul\(\s*-?\d+\s*,\s*-?\d+\s*\)|do\(\s*\)|don't\(\s*\))`)
	matches := re.FindAllString(s, -1)
	sum := 0
	for _, match := range matches {
		v := <-instruction
		if match[:3] == "mul" && v != "don't()" {
			x, y, err := parseMulNumbers(match)
			if err != nil {
				continue
			}
			sum += x * y
		}
		if match == "do()" || match == "don't()" {
			instruction <- match
		} else {
			instruction <- v
		}
	}
	return sum
}

func parseMulNumbers(s string) (int, int, error) {
	re := regexp.MustCompile(`mul\(\s*(-?\d+)\s*,\s*(-?\d+)\s*\)`)
	matches := re.FindStringSubmatch(s)

	if len(matches) != 3 {
		return 0, 0, fmt.Errorf("no valid 'mul(x,y)' found")
	}

	x, err1 := strconv.Atoi(matches[1])
	y, err2 := strconv.Atoi(matches[2])

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("error parsing numbers: %v, %v", err1, err2)
	}

	return x, y, nil
}

func TestRegValidate() {
	file, err := ScanIntoFIle("input.txt", 3)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	instruction := make(chan string, 2)
	instruction <- ""
	line_sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_sum += countValidMulSubstrings(line, instruction)
	}
	fmt.Println(line_sum)
	close(instruction)
}
