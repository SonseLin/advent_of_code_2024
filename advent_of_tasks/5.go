package adventoftasks

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func isExistingRule(rule, value int, rules map[int][]int) bool {
	for _, v := range rules[rule] {
		if v == value {
			return true
		}
	}
	return false
}

func middleNumber(numbers []int) int {
	return numbers[len(numbers)/2]
}

func isCorrectSequence(numbers []int, rules map[int][]int) bool {
	for i, number := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if !isExistingRule(number, numbers[j], rules) {
				return false
			}
		}
	}
	return true
}

func reorderSequence(sequence []int, rules map[int][]int) ([]int, error) {
	ordered := []int{}
	used := map[int]bool{}
	for len(ordered) < len(sequence) {
		for _, number := range sequence {
			if used[number] {
				continue
			}
			valid := true
			for _, numb := range sequence {
				if numb == number || used[numb] {
					continue
				}
				if !isExistingRule(number, numb, rules) {
					valid = false
					break
				}
			}
			if valid {
				ordered = append(ordered, number)
				used[number] = true
			}
		}
	}
	fmt.Println("sequence: ", sequence)
	fmt.Println("ordered :", ordered)
	return ordered, nil
}

func stringToIntSlice(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	var result []int
	for _, part := range parts {
		value, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func PrintQueue(mode string) {
	file, err := ScanIntoFIle("input.txt", 5)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rules := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	middle_sum_numbers := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) > 1 {
			rule, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting rule to int:", err)
				return
			}
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error converting value to int:", err)
				return
			}
			rules[rule] = append(rules[rule], value)
		} else if line != "" {
			numbers, err := stringToIntSlice(line)
			if err != nil {
				fmt.Println("Error converting line to int slice:", err)
				return
			}
			correct := isCorrectSequence(numbers, rules)
			if mode == "basic" {
				if correct {
					middle_sum_numbers += middleNumber(numbers)
				}
			} else if mode == "advanced" {
				if !correct {
					ordered, err := reorderSequence(numbers, rules)
					if err != nil {
						fmt.Println("Error reordering sequence:", err)
						return
					}
					middle_sum_numbers += middleNumber(ordered)
				}
			}
		}
	}
	fmt.Println(middle_sum_numbers)
}
