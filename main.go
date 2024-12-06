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
	case "2":
		adventoftasks.Task_2()
	case "2.1":
		adventoftasks.Task_2_ext()
	case "3":
		adventoftasks.TestRegValidate()
	case "4":
		adventoftasks.CeresSearch()
	case "5":
		adventoftasks.PrintQueue("basic")
	case "5.1":
		adventoftasks.PrintQueue("advanced")
	default:
	}
}
