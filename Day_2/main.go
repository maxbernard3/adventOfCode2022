package main

import (
	"bufio"
	"os"
)

func main() {
	result1 := puzzle1()
	println(result1)
	result2 := puzzle2()
	println(result2)
}

func puzzle1() int {
	totalScore := 0

	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			A := scanner.Text()[:1]
			B := scanner.Text()[2:]

			totalScore += victory(A, B) + playWorth(B)
		}
	}
	return totalScore
}

func puzzle2() int {
	totalScore := 0

	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			A := scanner.Text()[:1]
			B := scanner.Text()[2:]
			B = whatPlay(A, B)
			totalScore += victory(A, B) + playWorth(B)
		}
	}
	return totalScore
}

func victory(A string, B string) int {
	switch A {
	case "A":
		switch B {
		case "X":
			return 3
		case "Y":
			return 6
		default:
			return 0
		}

	case "B":
		switch B {
		case "X":
			return 0
		case "Y":
			return 3
		default:
			return 6
		}

	default:
		switch B {
		case "X":
			return 6
		case "Y":
			return 0
		default:
			return 3
		}
	}
}

func playWorth(B string) int {
	switch B {
	case "X":
		return 1
	case "Y":
		return 2
	default:
		return 3
	}
}

func whatPlay(A string, B string) string {
	switch A {
	case "A":
		switch B {
		case "X":
			return "Z"
		case "Y":
			return "X"
		default:
			return "Y"
		}

	case "B":
		switch B {
		case "X":
			return "X"
		case "Y":
			return "Y"
		default:
			return "Z"
		}

	default:
		switch B {
		case "X":
			return "Y"
		case "Y":
			return "Z"
		default:
			return "X"
		}
	}
}
