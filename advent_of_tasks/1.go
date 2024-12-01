package adventoftasks

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Task struct {
	list_1 []int
	list_2 []int
}

func (t *Task) fillTaskWithData(file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("incorrect format of line")
			continue
		}
		val1, err1 := strconv.Atoi(parts[0])
		if err1 != nil {
			fmt.Println(err1)
			continue

		}
		val2, err2 := strconv.Atoi(parts[1])
		if err2 != nil {
			fmt.Println(err2)
			continue
		}
		t.list_1 = append(t.list_1, val1)
		t.list_2 = append(t.list_2, val2)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func Task_1() {
	var t Task
	file, err := ScanIntoFIle("input.txt", 1)
	if err != nil {
		fmt.Println(err) // TODO: handle error
		return
	}
	defer file.Close()

	if err := t.fillTaskWithData(file); err != nil {
		fmt.Println(err)
		return
	}
	sort.Ints(t.list_1)
	sort.Ints(t.list_2)
	total_range := 0
	for i := 0; i < len(t.list_1); i++ {
		range_i := t.list_1[i] - t.list_2[i]
		if range_i < 0 {
			range_i = range_i * -1
		}
		total_range += range_i
	}
	fmt.Println(total_range)
}
