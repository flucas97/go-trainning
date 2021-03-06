package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := os.Args[1:]
	for _, number := range numbers {
		cast, _ := strconv.Atoi(number)
		fmt.Println(Solution(cast))
	}

}

func Solution(N int) int {
	binary := strconv.FormatInt(int64(N), 2)
	binaryArray := strings.Split(binary, "")
	zerosCount, maxZerosCount := 0, 0

	for _, binaryNumber := range binaryArray {
		if binaryNumber == "1" {
			if zerosCount > maxZerosCount {
				maxZerosCount = zerosCount
			}
			zerosCount = 0

		} else {
			zerosCount++
		}
	}
	return maxZerosCount
}
