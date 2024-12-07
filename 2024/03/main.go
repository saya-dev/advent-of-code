package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getSuffix (s string) string {
	if s == "" {
		return ""
	}
	return s[len(s)-1:]
}

func isDigit (s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return true
	}
	return false
}

func main()  {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	
	enabled := true
	instruction := ""
	calc := ""
	sum := 0
	enabledSum := 0

	for scanner.Scan() {
		r := scanner.Text()

		if instructionSuffix := getSuffix(instruction); enabled {
			if 
			(instruction == "" && r == "d") || (instruction == "d" && r == "o") || (instructionSuffix == "o" && r == "n" ) || 
			(instructionSuffix == "n" && r == "'" ) || (instructionSuffix == "'" && r == "t") ||(instructionSuffix == "t" && r == "(" ) {
				instruction += r
				continue
			} else if instructionSuffix == "(" && r == ")" {
				enabled = false
				instruction = ""
				continue
			} else {
				instruction = ""
			}
		} else {
			if (instruction == "" && r == "d") || (instruction == "d" && r == "o" ) || (instructionSuffix == "o" && r == "(") {
				instruction += r
				continue
			} else if instructionSuffix == "(" && r == ")" {
				enabled = true
				instruction = ""
				continue
			} else {
				instruction = ""
			}
		}

		if calcSuffix := getSuffix(calc); (calc == "" && r != "m") || (calc == "m" && r!= "u") || (calcSuffix == "u" && r != "l") || 
		(calcSuffix == "l" && r != "(") || (calcSuffix == "(" && !isDigit(r)) || (isDigit(calcSuffix) && !isDigit(r) && r != "," && r!= ")") {
			calc = ""
			continue
		}

		calc += r

		if r == ")" {
			x, y, _ := strings.Cut(strings.Trim(calc, "mul()"), ",")
			ix, _ := strconv.Atoi(x)
			iy, _ := strconv.Atoi(y)

			sum += ix * iy
			calc = ""

			if enabled {
				enabledSum += ix * iy
			}
		}
	}

	fmt.Println("Result:", sum)
	fmt.Println("Enabled Result:", enabledSum)
}