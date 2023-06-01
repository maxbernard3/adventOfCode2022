package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	dirArr := loadDir()

	result1 := puzzle1(dirArr)
	println(result1)

	result2 := puzzle2(dirArr)
	println(result2)
}

type Dir struct {
	name   string
	parent *Dir
	subDir []*Dir
	size   int
}

func puzzle1(dirArr []*Dir) int {
	counter := 0
	for _, a := range dirArr {
		if a.size <= 100000 {
			counter += a.size
		}
	}
	return counter
}

func puzzle2(dirArr []*Dir) int {
	minSize := dirArr[0].size + 30000000 - 70000000

	dirArr = quicksort(dirArr, 0, len(dirArr)-1)

	for _, a := range dirArr {
		if a.size > minSize {
			return a.size
		}
	}

	return 0
}

func loadDir() []*Dir {
	var dirArr []*Dir

	file, err := os.Open("in.txt")
	defer file.Close()
	if err != nil {
		println(err)
	} else {
		scanner := bufio.NewScanner(file)
		var curent *Dir
		for scanner.Scan() {
			if scanner.Text()[0] == "$"[0] {
				if scanner.Text() == "$ cd .." {
					curent = curent.parent
				} else if scanner.Text()[:4] == "$ cd" {
					a := true
					if curent != nil {
						for _, child := range curent.subDir {
							if child.name == scanner.Text()[5:] && a {
								curent = child
								a = false
							}
						}
					}
					if a {
						curent = &Dir{scanner.Text()[5:], nil, nil, 0}
						dirArr = append(dirArr, curent)
					}
				}
			}
			if scanner.Text()[:3] == "dir" {
				temp := &Dir{scanner.Text()[4:], curent, nil, 0}
				dirArr = append(dirArr, temp)
				curent.subDir = append(curent.subDir, temp)
			}
			line := strings.Fields(scanner.Text())
			if res, err := strconv.Atoi(line[0]); err == nil {
				curent.size += res
			}
		}
	}

	for i := len(dirArr) - 1; i >= 0; i-- {
		for _, a := range dirArr[i].subDir {
			dirArr[i].size += a.size
		}
	}

	return dirArr
}

func quicksort(slice []*Dir, min int, max int) []*Dir {
	if min <= max {
		pi := partition(slice, min, max)

		quicksort(slice, min, pi-1)
		quicksort(slice, pi+1, max)
	}

	return slice
}

func partition(slice []*Dir, min int, max int) int {
	pivot := slice[max].size
	i := min - 1

	for j := min; j < max; j++ {
		if slice[j].size < pivot {
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
