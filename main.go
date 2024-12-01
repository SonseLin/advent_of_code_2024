package main

import (
	adventoftasks "adventOfGo2024/advent_of_tasks"
	"fmt"
)

func main() {
	var val string
	fmt.Scan(&val)
	switch val {
	case "1":
		adventoftasks.Task_1() // 任务1
	case "1.1":
		adventoftasks.Task_1_ext()
	default:
	}
}
