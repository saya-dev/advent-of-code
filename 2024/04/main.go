package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 137 - 2

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum1 := 0
	sum2 := 0
	rows := [][]string{}
	index := 0

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		rows = append(rows, row)

		for i := 0; i < len(row); i++ {
			r := row[i]

			if r == "A" {
				continue
			}

			// horizontal
			if r == "X" && i < 137 && row[i+1] == "M" && row[i+2] == "A" && row[i+3] == "S" {
				sum1 += 1
			}
			// horizontal reverse
			if r == "S" && i < 137 && row[i+1] == "A" && row[i+2] == "M" && row[i+3] == "X" {
				sum1 += 1
			}
			// vertical
			if r == "X" && index > 2 && rows[index-1][i] == "M" && rows[index-2][i] == "A" && rows[index-3][i] == "S" {
				sum1 += 1
			}
			// vertical reverse
			if r == "S" && index > 2 && rows[index-1][i] == "A" && rows[index-2][i] == "M" && rows[index-3][i] == "X" {
				sum1 += 1
			}
			// diagonal left
			if r == "X" && index > 2 && i > 2 && rows[index-1][i-1] == "M" && rows[index-2][i-2] == "A" && rows[index-3][i-3] == "S" {
				sum1 += 1
			}
			// diagonal left reverse
			if r == "S" && index > 2 && i > 2 && rows[index-1][i-1] == "A" && rows[index-2][i-2] == "M" && rows[index-3][i-3] == "X" {
				sum1 += 1
			}
			// diagonal right
			if r == "X" && index > 2 && i < 137 && rows[index-1][i+1] == "M" && rows[index-2][i+2] == "A" && rows[index-3][i+3] == "S" {
				sum1 += 1
			}
			// diagonal right reverse
			if r == "S" && index > 2 && i < 137 && rows[index-1][i+1] == "A" && rows[index-2][i+2] == "M" && rows[index-3][i+3] == "X" {
				sum1 += 1
			}
			// x-mas
			if (r == "M" || r == "S") && index > 1 && i < 138 && rows[index-1][i+1] == "A" && ((r == "M" && rows[index-2][i+2] == "S") || (r == "S" && rows[index-2][i+2] == "M")) &&
				(row[i+2] == "M" || row[i+2] == "S") && ((row[i+2] == "M" && rows[index-2][i] == "S") || (row[i+2] == "S" && rows[index-2][i] == "M")) {
				sum2 += 1
			}
		}

		index += 1
	}

	fmt.Println("XMAS:", sum1)
	fmt.Println("X-MAS:", sum2)
}
