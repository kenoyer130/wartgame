package models

import (
	"log"
	"strconv"
)

// TODO add generic helper function that returns an initilized 2d array of x,y size
// Note this is not an actual matrix as x and y can differ in size

func New2DArray[V any](x int, y int) [][]V {

	m := make([][]V, x)

	for i := range m {
		m[i] = make([]V, y)
	}

	return m
}

func Print2DArray (grid [][]int) {	

	row := ""
	for c := 0; c < len(grid); c++ {
		for r := 0; r <len(grid[c]); r ++ {	
			row += strconv.Itoa(grid[c][r])
		}			
		log.Println(row)
		row = ""
	}
}
