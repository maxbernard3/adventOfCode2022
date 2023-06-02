package main

import (
	"bufio"
	"os"
)

//262752, 293280

func main() {
	inList := loadList()

	result1 := puzzle1(inList)
	println(result1)

	result2 := puzzle2(inList)
	println(result2)
}

func puzzle1(inList [][]int) int {
	var counter int

	for x, row := range inList {
		for y := range row {
			if x == 0 || y == 0 || x == len(inList)-1 || y == len(inList[0])-1 || !look(inList, x, y) {
				counter++
			}
		}
	}

	return counter
}

func puzzle2(inList [][]int) int {
	var viewArr []int

	for x, row := range inList {
		for y := range row {
			if x == 0 || y == 0 || x == len(inList)-1 || y == len(inList[0])-1 {
				viewArr = append(viewArr, 0)
			} else {
				viewArr = append(viewArr, senicScore(inList, x, y))
			}
		}
	}

	counter := quicksort(viewArr, 0, len(viewArr)-1)[len(viewArr)-1]

	return counter
}

func loadList() [][]int {
	var fList [][]int
	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var tempList []int
			for _, i := range scanner.Text() {
				tempList = append(tempList, int(i))
			}
			fList = append(fList, tempList)
		}
	}

	return fList
}

func look(list [][]int, x int, y int) bool {
	lx := len(list)
	ly := len(list[0])
	r := 0

	for i := x + 1; i < lx; i++ {
		if list[x][y] <= list[i][y] {
			r++
			i = lx
		}
	}
	for i := x - 1; i >= 0; i-- {
		if list[x][y] <= list[i][y] {
			r++
			i = -1
		}
	}
	for i := y + 1; i < ly; i++ {
		if list[x][y] <= list[x][i] {
			r++
			i = ly
		}
	}
	for i := y - 1; i >= 0; i-- {
		if list[x][y] <= list[x][i] {
			r++
			i = -1
		}
	}
	if r < 4 {
		return false
	} else {
		return true
	}
}

func senicScore(list [][]int, x int, y int) int {
	lx := len(list)
	ly := len(list[0])
	up := lx - x - 1
	down := x
	left := ly - y - 1
	right := y

	for i := x + 1; i < lx; i++ {
		if list[x][y] <= list[i][y] {
			up = i - x
			i = lx
		}
	}
	for i := x - 1; i >= 0; i-- {
		if list[x][y] <= list[i][y] {
			down = x - i
			i = -1
		}
	}
	for i := y + 1; i < ly; i++ {
		if list[x][y] <= list[x][i] {
			left = i - y
			i = ly
		}
	}
	for i := y - 1; i >= 0; i-- {
		if list[x][y] <= list[x][i] {
			right = y - i
			i = -1
		}
	}

	return up * down * left * right
}

func quicksort(slice []int, min int, max int) []int {
	if min <= max {
		pi := partition(slice, min, max)

		quicksort(slice, min, pi-1)
		quicksort(slice, pi+1, max)
	}

	return slice
}

func partition(slice []int, min int, max int) int {
	pivot := slice[max]
	i := min - 1

	for j := min; j < max; j++ {
		if slice[j] < pivot {
			i++
			tmp2 := slice[i]
			slice[i] = slice[j]
			slice[j] = tmp2
		}
	}
	i++
	tmp := slice[i]
	slice[i] = slice[max]
	slice[max] = tmp
	return i
}
