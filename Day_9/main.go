package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func main() {
	inList := loadList()

	result1 := puzzle1(inList)
	println(result1)

	result2 := puzzle2(inList)
	println(result2)
}

func puzzle1(inList [][2]int) int {
	H := [2]int{0, 0}
	T := [2]int{0, 0}

	var posHistory [][2]int

	for _, tup := range inList {
		for i := 0; i < tup[1]; i++ {
			posHistory = checkHistory(posHistory, T)

			switch tup[0] {
			case 0:
				H[0]++
			case 1:
				H[0]--
			case 2:
				H[1]++
			case 3:
				H[1]--
			}

			T = stepT(T, H)
		}
	}

	return len(posHistory)
}

func checkHistory(posHistory [][2]int, T [2]int) [][2]int {
	revisit := false
	for _, pos := range posHistory {
		if pos == T {
			revisit = true
		}
	}

	if !revisit {
		posHistory = append(posHistory, T)
	}

	return posHistory
}

func stepT(T [2]int, H [2]int) [2]int {
	if math.Abs(float64(T[0]-H[0])) == 2 && T[1] == H[1] {
		T[0] += Normalize(H[0] - T[0])
	} else if math.Abs(float64(T[1]-H[1])) == 2 && T[0] == H[0] {
		T[1] += Normalize(H[1] - T[1])
	} else if math.Abs(float64(T[1]-H[1])) >= 2 && math.Abs(float64(T[0]-H[0])) >= 1 {
		T[1] += Normalize(H[1] - T[1])
		T[0] += Normalize(H[0] - T[0])
	} else if math.Abs(float64(T[1]-H[1])) >= 1 && math.Abs(float64(T[0]-H[0])) >= 2 {
		T[1] += Normalize(H[1] - T[1])
		T[0] += Normalize(H[0] - T[0])
	}

	return T
}

func Normalize(a int) int {
	if math.Signbit(float64(a)) {
		return -1
	}
	return 1
}

func puzzle2(inList [][2]int) int {
	var T [10][2]int

	var posHistory [][2]int

	for _, tup := range inList {
		for i := 0; i < tup[1]; i++ {
			posHistory = checkHistory(posHistory, T[9])

			switch tup[0] {
			case 0:
				T[0][0]++
			case 1:
				T[0][0]--
			case 2:
				T[0][1]++
			case 3:
				T[0][1]--
			}

			for j := 1; j < 10; j++ {
				T[j] = stepT(T[j], T[j-1])
			}
		}
	}

	return len(posHistory)
}

func loadList() [][2]int {
	var fList [][2]int
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var tempList [2]int

			switch scanner.Text()[:1] {
			case "U":
				tempList[0] = 0
			case "D":
				tempList[0] = 1
			case "R":
				tempList[0] = 2
			case "L":
				tempList[0] = 3
			}
			d, _ := strconv.Atoi(scanner.Text()[2:])
			tempList[1] = d
			fList = append(fList, tempList)
		}
	}

	return fList
}
