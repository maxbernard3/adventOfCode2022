package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func main() {
	result1 := puzzle1()
	println(result1)
	result2 := puzzle2()
	println(result2)
}

func puzzle1() int {
	res := 0
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			a := getInt(scanner.Text())
			if checkOverlaps(a) {
				res++
			}
		}
	}

	return (res)
}

func puzzle2() int {
	res := 0
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			a := getInt(scanner.Text())
			if checkAnyOverlaps(a) {
				res++
			}
		}
	}

	return (res)
}

func getInt(in string) [4]int {
	var strRes [4]string
	j := 0
	for i := 0; i < len(in); i++ {
		if unicode.IsNumber(rune(in[i])) {
			strRes[j] += string(in[i])
		} else {
			j++
		}
	}

	var res [4]int
	for i := 0; i < 4; i++ {
		res[i], _ = strconv.Atoi(strRes[i])
	}
	return res
}

func checkOverlaps(in [4]int) bool {
	if (in[0] >= in[2] && in[1] <= in[3]) || (in[2] >= in[0] && in[3] <= in[1]) {
		return true
	} else {
		return false
	}
}

func checkAnyOverlaps(in [4]int) bool {
	if in[1] < in[2] || in[0] > in[3] {
		return false
	} else {
		return true
	}
}
