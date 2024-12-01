package adventoftasks

import (
	"fmt"
)

func Task_1_ext() {
	var t Task
	file, err := ScanIntoFIle("input.txt", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // close file at the end of the program
	if err := t.fillTaskWithData(file); err != nil {
		fmt.Println(err)
		return
	}
	total_range := 0
	entries := make(map[int]int)
	entries_2 := make(map[int]int)
	for _, num := range t.list_1 {
		entries[num]++
	}
	for _, num := range t.list_2 {
		entries_2[num]++
	}
	for num, count := range entries {
		if entries_2[num] > 0 {
			total_range += count * entries_2[num] * num
		}
	}
	fmt.Println(total_range)
}
