package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diff(a, b int) int {
   if a < b {
      return b - a
   }
   return a - b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list1 := []int{}
	list2 := []int{}
	
	for scanner.Scan() {
		first, second, _ := strings.Cut(scanner.Text(), "   ")
		
		firstInt, err := strconv.Atoi(first)
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, firstInt)

		secondInt, err := strconv.Atoi(second)
		if err != nil {
			log.Fatal(err)
		}
		list2 = append(list2, secondInt)
	}

	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))

	distance := 0
	similarity := 0

	for x := 0; x < len(list1); x++ {
		distance += diff(list1[x], list2[x])

		for y := 0; y < len(list2); y++ {
			if list1[x] == list2[y] {
				similarity += list1[x]
			}
		}
	}

	fmt.Println("Distance:", distance)
	fmt.Println("Similarity:", similarity)
}