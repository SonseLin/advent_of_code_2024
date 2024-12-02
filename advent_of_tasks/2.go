package adventoftasks

import (
	"bufio"
	"fmt"
	"strings"
)

func isDecreaseOrIncreaseOnly(reports []int) bool {
	decreasing_values := 0
	increasing_values := 0
	for i := 0; i < len(reports)-1; i++ {
		if reports[i] > reports[i+1] && reports[i]-3 <= reports[i+1] {
			decreasing_values++
		} else if reports[i] < reports[i+1] && reports[i]+3 >= reports[i+1] {
			increasing_values++
		} else {
			return false
		}
		if decreasing_values > 0 && increasing_values > 0 {
			return false
		}
	}
	return true
}

func is_safe_report(reports []int) int {
	for i, report := range reports {
		if i < 1 {
			continue
		}
		if report-2 > reports[i] {
			return 0
		} else if report+2 < reports[i] {
			return 0
		}
	}
	return 1
}

func Task_2() {
	file, err := ScanIntoFIle("input.txt", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe_reports := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		numbers_int, err := stringsToInts(numbers)
		if err != nil {
			fmt.Println(err)
			return
		}
		if isDecreaseOrIncreaseOnly(numbers_int) {
			safe_reports += is_safe_report(numbers_int)
		}
	}
	fmt.Println(safe_reports)
}
