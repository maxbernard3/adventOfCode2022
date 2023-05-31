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
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			for i := 0; i < len(scanner.Text())-4; i++ {
				allunique := true
				for a, chA := range scanner.Text()[i : i+4] {
					for b, chB := range scanner.Text()[i : i+4] {
						if chA == chB && a != b {
							allunique = false
						}
					}
				}
				if allunique {
					return i + 4
				}
			}
		}
	}
	return 0
}

func puzzle2() int {
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			for i := 0; i < len(scanner.Text())-14; i++ {
				allunique := true
				for a, chA := range scanner.Text()[i : i+14] {
					for b, chB := range scanner.Text()[i : i+14] {
						if chA == chB && a != b {
							allunique = false
						}
					}
				}
				if allunique {
					return i + 14
				}
			}
		}
	}
	return 0
}
