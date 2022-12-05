package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
	result1 := puzzle1()
	println(result1)
	result2 := puzzle2()
	println(result2)
}

func puzzle1() int {
	var res = 0

	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c := getCompartment(scanner.Text())
			d := compareCompartment(c[0], c[1])
			res += getPriority(d)
		}
	}
	return res
}

func puzzle2() int {
	var res = 0
	i := 0
	var get3 [3]string

	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			get3[i] = scanner.Text()
			i++
			if i == 3 {
				i = 0
				r := compareRucksacks(get3[0], get3[1], get3[2])
				res += getPriority(string(r))
			}
		}
	}
	return res
}

func getCompartment(in string) []string {
	len := int(len(in) / 2)
	return ([]string{in[:len], in[len:]})
}

func compareCompartment(A string, B string) string {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				return string(A[i])
			}
		}
	}
	return ""
}

func getPriority(in string) int {
	byteIn := in[0]
	out := 0
	if unicode.IsLower(rune(byteIn)) {
		for i := 0; i < 26; i++ {
			if in == string(alphabet[i]) {
				out = i + 1
				i = 50
			}
		}
	} else {
		for i := 0; i < 26; i++ {
			if in == strings.ToUpper(string(alphabet[i])) {
				out = i + 27
				i = 50
			}
		}
	}
	return out
}

func compareRucksacks(A string, B string, C string) byte {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			for k := 0; k < len(C); k++ {
				if A[i] == B[j] && B[j] == C[k] {
					return A[i]
				}
			}
		}
	}
	return 0
}
