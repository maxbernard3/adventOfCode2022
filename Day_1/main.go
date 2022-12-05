package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	result1 := puzzle1()
	println(result1)
	result2 := puzzle2()
	println(result2)
}

func puzzle1() int {
	bigfood := 0
	curentfood := 0

	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			a, _ := strconv.Atoi(scanner.Text())
			if a == 0 {
				curentfood = 0
			} else {
				curentfood += a
				if curentfood > bigfood {
					bigfood = curentfood
				}
			}
		}
	}
	return (bigfood)
}

func puzzle2() int {
	var arrFood [1000]int
	i := 0
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			a, _ := strconv.Atoi(scanner.Text())
			if a != 0 {
				arrFood[i] += a
			} else {
				i++
			}
		}
	}

	arrSlice := arrFood[:]
	sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))
	res := arrFood[0] + arrFood[1] + arrFood[2]

	return (res)
}
