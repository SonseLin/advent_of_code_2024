package adventoftasks

import (
	"fmt"
	"os"
	"strconv"
)

func ScanIntoFIle(filepath string, day int) (*os.File, error) {
	return os.Open(fmt.Sprintf("advent_of_tasks/src/%d/%s", day, filepath))
}

func stringsToInts(strings []string) ([]int, error) {
	ints := make([]int, 0, len(strings))
	for _, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования строки '%s': %v", str, err)
		}
		ints = append(ints, num)
	}
	return ints, nil
}
