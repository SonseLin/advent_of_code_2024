package adventoftasks

import (
	"bufio"
	"fmt"
)

func CeresSearch() {
	file, err := ScanIntoFIle("input.txt", 4)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	matrix := make([][]byte, 100000)
	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		matrix[i] = make([]byte, len(line))
		for j, char := range line {
			matrix[i][j] = byte(char)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	word := "XMAS"
	counts := countWordOccurrences(matrix[:i], word)
	printCounts(counts, word)
	total := 0
	for _, v := range counts {
		total += v
	}
	fmt.Println("Total:", total)
}

func countWordOccurrences(matrix [][]byte, word string) map[string]int {
	counts := map[string]int{
		"Horizontal Right":    0,
		"Horizontal Left":     0,
		"Vertical Down":       0,
		"Vertical Up":         0,
		"Diagonal Down-Right": 0,
		"Diagonal Up-Left":    0,
		"Diagonal Down-Left":  0,
		"Diagonal Up-Right":   0,
	}

	rows := len(matrix)
	cols := len(matrix[0])
	wordBytes := []byte(word)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == wordBytes[0] {
				if checkDirection(matrix, r, c, wordBytes, 0, 1) {
					counts["Horizontal Right"]++
				}
				if checkDirection(matrix, r, c, wordBytes, 0, -1) {
					counts["Horizontal Left"]++
				}
				if checkDirection(matrix, r, c, wordBytes, 1, 0) {
					counts["Vertical Down"]++
				}
				if checkDirection(matrix, r, c, wordBytes, -1, 0) {
					counts["Vertical Up"]++
				}
				if checkDirection(matrix, r, c, wordBytes, 1, 1) {
					counts["Diagonal Down-Right"]++
				}
				if checkDirection(matrix, r, c, wordBytes, -1, -1) {
					counts["Diagonal Up-Left"]++
				}
				if checkDirection(matrix, r, c, wordBytes, 1, -1) {
					counts["Diagonal Down-Left"]++
				}
				if checkDirection(matrix, r, c, wordBytes, -1, 1) {
					counts["Diagonal Up-Right"]++
				}
			}
		}
	}
	return counts
}

func checkDirection(matrix [][]byte, startRow, startCol int, word []byte, rowDir, colDir int) bool {
	for i := 0; i < len(word); i++ {
		newRow := startRow + i*rowDir
		newCol := startCol + i*colDir
		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) || matrix[newRow][newCol] != word[i] {
			return false
		}
	}
	return true
}

func printCounts(counts map[string]int, word string) {
	fmt.Printf("Occurrences of the word '%s':\n", word)
	for direction, count := range counts {
		fmt.Printf("%s: %d\n", direction, count)
	}
}
