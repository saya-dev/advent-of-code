package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func diff(a, b int) int {
   if a < b {
      return b - a
   }
   return a - b
}

func main()  {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0
	fixableReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		levels := []int{}

		for {
			x, rest, found := strings.Cut(line, " ")

			ix, _ := strconv.Atoi(x)
			levels = append(levels, ix)

			if !found {
				break
			}
			
			line = rest
		}
		
		increasing := levels[0] < levels[1]
		errors := 0

		for i := 0; i < len(levels) - 1; i++ {
			difference := diff(levels[i], levels[i+1])

			if increasing != (levels[i] < levels[i+1]) || difference < 1 || difference > 3  {
				if errors == 2 {
					break	
				}
				errors += 1
			}
		}

		if errors == 0 {
			safeReports += 1
		} else if errors == 1 {
			fixableReports += 1
		}
	}

	fmt.Println("Safe Reports:", safeReports)
	fmt.Println("Fixable Reports:", fixableReports)
	fmt.Println("Safe & Fixable Reports:", safeReports + fixableReports)
}