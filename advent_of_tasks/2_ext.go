package adventoftasks

import (
	"bufio"
	"fmt"
	"strings"
)

func validateOneFailure(reports []int) bool {
	for i := 0; i < len(reports); i++ {
		data1 := append([]int{}, reports[:i]...)
		data1 = append(data1, reports[i+1:]...)
		if safeReportsValidate(data1) {
			return true
		}
	}
	return false
}

func safeReportsValidate(data []int) bool {
	return isDecreaseOrIncreaseOnly(data)
}

func Task_2_ext() {
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
		if safeReportsValidate(numbers_int) || validateOneFailure(numbers_int) {
			safe_reports++
		}
	}
	fmt.Println(safe_reports)
}
