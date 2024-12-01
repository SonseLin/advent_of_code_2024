package adventoftasks

import (
	"fmt"
	"os"
)

func ScanIntoFIle(filepath string, day int) (*os.File, error) {
	return os.Open(fmt.Sprintf("advent_of_tasks/src/%d/%s", day, filepath))
}
