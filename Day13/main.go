package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	input = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`
)

type tuple struct {
	x, y int
}

type fold struct {
	where   string
	howMuch int
}

func main() {
	s := strings.Split(input, "\n")
	var dots []tuple
	var folds []fold
	var largestX = 0
	var largestY = 0
	//var table [][]int

	for _, line := range s {
		if strings.Contains(line, "fold along ") {
			splitLine := strings.Split(line, "=")
			howMuch, _ := strconv.Atoi(splitLine[1])
			where := splitLine[0][len(splitLine[0])-1:]
			folds = append(folds, fold{where, howMuch})
		} else if strings.Contains(line, ",") {
			f := strings.Split(line, ",")
			x, _ := strconv.Atoi(f[0])
			y, _ := strconv.Atoi(f[1])
			dots = append(dots, tuple{x, y})
			if largestX < x {
				largestX = x
			}
			if largestY < y {
				largestY = y
			}
		}
	}
	matrix := make([][]string, largestY+1)
	for i := range matrix {
		matrix[i] = make([]string, largestX+1)
	}

	for l := 0; l < largestX+1; l++ {
		for i := 0; i < largestY+1; i++ {
			matrix[i][l] = "."
		}
	}
	addLocationsToMatrix(matrix, dots)
	printMatrix(matrix)
	fmt.Println(folds)
}


func addLocationsToMatrix(givenMatrix [][]string, dots []tuple) {
	for _, dot := range dots {
		givenMatrix[dot.y][dot.x] = "#"
	}
}

func printMatrix(m [][]string) {
	for i, i2 := range m {
		fmt.Print(i)
		fmt.Println(i2)
	}
}
