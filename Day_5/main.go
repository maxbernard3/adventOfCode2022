package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	/*a := readInstruction("move 2 from 2 to 1")
	b := crateToArr()
	res := moveCrate(b, a)
	for i := 0; i < 9; i++ {
		println(b[i][0])
		println("-------")
		println(b[i][1])
	}*/
	a := readInstruction("move 2 from 2 to 1")
	b := crateToArr()
	moveCrate(b, a)
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

		}
	}

	return (res)
}

func crateToArr() [100][9]string {
	var strRes [100][9]string
	j := 0
	file, err := os.Open("crate.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			scanner.Text()
			k := 0
			for i := 1; k < 9; i += 4 {
				strRes[j][k] = string(scanner.Text()[i])
				k++
			}
			j++
		}
	}

	return strRes
}

func readInstruction(in string) [3]int {
	re := regexp.MustCompile("[0-9]+")
	a := (re.FindAllString(in, -1))

	var res [3]int
	for i := 0; i < len(a); i++ {
		res[i], _ = strconv.Atoi(a[i])
	}

	return res
}

func moveCrate(crate [100][9]string, in [3]int) [100][9]string {
	tempCrate := crate
	maxout := 0
	maxin := 0
	for i := 0; i < 99; i++ {
		if crate[i][in[1]] == "" {
			maxout = i
		}
		if crate[i][in[2]] == "" {
			maxin = i
		}
	}
	println(maxout, maxin)
	return tempCrate
}
