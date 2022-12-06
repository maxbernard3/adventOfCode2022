package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	result1 := puzzle1()
	println(result1)
	result2 := puzzle2()
	println(result2)
}

func puzzle1() string {
	crate := crateToArr()
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			inst := readInstruction(scanner.Text())
			crate = moveCrate(crate, inst)
		}
	}

	res := topCrate(crate)

	return (res)
}

func puzzle2() string {
	crate := crateToArr()
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			inst := readInstruction(scanner.Text())
			crate = moveCrateStacks(crate, inst)
		}
	}

	res := topCrate(crate)

	return (res)
}

func crateToArr() [9][]string {
	var strRes [9][]string
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
				if string(scanner.Text()[i]) != " " {
					strRes[k] = append(strRes[k], string(scanner.Text()[i]))
				}
				k++
			}
		}
	}
	for i := 0; i < 9; i++ {
		strRes[i] = reverse(strRes[i])
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
	res[1] = res[1] - 1
	res[2] = res[2] - 1

	return res
}

func moveCrate(crate [9][]string, in [3]int) [9][]string {

	for i := 0; i < in[0]; i++ {
		lenRm := len(crate[in[1]]) - 1
		crate[in[2]] = append(crate[in[2]], crate[in[1]][lenRm])
		crate[in[1]] = crate[in[1]][:lenRm]
	}

	return crate
}

func moveCrateStacks(crate [9][]string, in [3]int) [9][]string {

	lenRm := len(crate[in[1]]) - in[0]

	tempCrate := crate[in[1]][lenRm:]

	for i := 0; i < len(tempCrate); i++ {
		crate[in[2]] = append(crate[in[2]], tempCrate[i])
	}
	crate[in[1]] = crate[in[1]][:lenRm]

	return crate
}

func topCrate(crate [9][]string) string {
	res := ""
	for i := 0; i < 9; i++ {
		res += crate[i][len(crate[i])-1]
	}
	return res
}

func reverse(input []string) []string {
	var output []string
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}
