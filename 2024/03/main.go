package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	sum := 0
	enabledSum := 0
	text := ""
	text2 := ""
	enabled := true

	for scanner.Scan() {
		r := scanner.Text()

		if enabled {
			if text2 == "" && r == "d" {
				text2 += r
			} else if strings.HasSuffix(text2,"d") && r == "o" {
				text2 += r
			} else if strings.HasSuffix(text2,"o") && r == "n" {
				text2 += r
			} else if strings.HasSuffix(text2,"n") && r == "'" {
				text2 += r;
			} else if strings.HasSuffix(text2,"'") && r == "t" {
				text2 += r;
			} else if strings.HasSuffix(text2,"t") && r == "(" {
				text2 += r
			} else if strings.HasSuffix(text2,"(") && r == ")" {
				enabled = false
				text2 = ""
			} else {
				text2 = ""
			}
		} else {
			if text2 == "" && r == "d" {
				text2 += r
			} else if strings.HasSuffix(text2,"d") && r == "o" {
				text2 += r
			} else if strings.HasSuffix(text2,"o") && r == "(" {
				text2 += r
			} else if strings.HasSuffix(text2,"(") && r == ")" {
				enabled = true
				text2 = ""
			} else {
				text2 = ""
			}
		}

		if text == "" && r != "m" {
			continue
		}

		if text == "m" && r!= "u" {
			text = ""
			continue
		}

		if strings.HasSuffix(text, "u") && r != "l" {
			text = ""
			continue
		}

		if strings.HasSuffix(text, "l") && r != "(" {
			text = ""
			continue
		}

		if strings.HasSuffix(text, "(") && !isDigit(r){
			text = ""
			continue
		}

		if text != "" && isDigit(text[len(text)-1:]) && !isDigit(r) && r != "," && r!= ")" {
			text = ""
			continue
		}

		text += r

		if strings.HasSuffix(text, ")") {
			text = strings.TrimPrefix(text, "mul(")
			text = strings.TrimSuffix(text, ")")

			x, y, _ := strings.Cut(text, ",")

			ix, _ := strconv.Atoi(x)
			iy, _ := strconv.Atoi(y)

			sum += ix * iy

			if enabled {
				enabledSum += ix * iy
			}

			text = ""
		}
	}


	fmt.Println("Result:", sum)
	fmt.Println("Enabled Result:", enabledSum)
}