package main

import (
	"bufio"
	"fmt"
	"github.com/TwiN/go-color"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	columnIndex int
	rowIndex    int
}

type matrices struct {
	inputMatrix   [][]int
	visitedMatrix [][]int
}

func (matrix *matrices) initialInputMatrix(height, width int) {
	matrix.inputMatrix = make([][]int, height)

	for k := range matrix.inputMatrix {
		matrix.inputMatrix[k] = make([]int, width)
	}
}

func (matrix *matrices) initialVisitedMatrix(height, width int) {
	matrix.visitedMatrix = make([][]int, height)

	for k := range matrix.visitedMatrix {
		matrix.visitedMatrix[k] = make([]int, width)
	}
}

func (matrix *matrices) insertRandValue(colour int) {
	for _, row := range matrix.inputMatrix {
		for k := range row {
			row[k] = rand.Intn(colour)
		}
	}
}

func (matrix *matrices) fillUpEmptyMatrix(colour int) {
	for _, row := range matrix.visitedMatrix {
		for k := range row {
			row[k] = (colour - colour) - 1
		}
	}
}

func (matrix *matrices) printMatrix(maxResult []coordinate) {
	for i := 0; i < len(matrix.inputMatrix); i++ {
		var strLine strings.Builder

	loop:
		for j := 0; j < len(matrix.inputMatrix[i]); j++ {
			for k := range maxResult {
				cor := maxResult[k]

				if cor.columnIndex == i && cor.rowIndex == j {
					strLine.WriteString(color.InRed(strconv.Itoa(matrix.inputMatrix[i][j]) + " "))
					continue loop
				}
			}

			strLine.WriteString(strconv.Itoa(matrix.inputMatrix[i][j]) + " ")
		}

		fmt.Println(strLine.String())
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			fmt.Println(err)
		}
	}(out)

	var height, width, colour int
	_, err := fmt.Fscan(in, &height, &width, &colour)
	if err != nil {
		fmt.Println(err)
		return
	}

	var matrix matrices
	matrix.initialInputMatrix(height, width)
	matrix.initialVisitedMatrix(height, width)

	matrix.insertRandValue(colour)
	matrix.fillUpEmptyMatrix(colour)

	maxResult := searchMatches(height, width, colour, matrix)

	fmt.Println(" ")
	matrix.printMatrix(maxResult)
	fmt.Println(" ")
	fmt.Println(maxResult)
}

func searchMatches(height, width, colour int, matrix matrices) []coordinate {
	var maxResult = make([]coordinate, 0, height*width)
	for i := 0; i < colour; i++ {
		matrix.fillUpEmptyMatrix(colour)

		for iC, row := range matrix.inputMatrix {
			for iR, val := range row {
				if val != i {
					if matrix.visitedMatrix[iC][iR] == i+1 {
						continue
					}
					matrix.visitedMatrix[iC][iR] = i + 1
				} else {
					tempResult := make([]coordinate, 0, height*width)
					tempResult = deepSearch(iC, iR, i, tempResult, matrix)
					if len(tempResult) > len(maxResult) {
						maxResult = tempResult
					}
				}
			}
		}
	}

	return maxResult
}

func deepSearch(iC, iR, i int, tempResult []coordinate, matrix matrices) []coordinate {
	if iC < 0 || iC >= len(matrix.inputMatrix) || iR < 0 || iR >= len(matrix.inputMatrix[0]) { //out of range
		return tempResult
	}
	if matrix.inputMatrix[iC][iR] != i { //not exist
		return tempResult
	}

	//already visited
	if matrix.visitedMatrix[iC][iR] == i || matrix.visitedMatrix[iC][iR] == i+1 {
		return tempResult
	}

	matrix.visitedMatrix[iC][iR] = i
	if matrix.inputMatrix[iC][iR] == i {
		tempResult = append(tempResult, coordinate{iC, iR})

		tempResult = deepSearch(iC-1, iR, i, tempResult, matrix)
		tempResult = deepSearch(iC+1, iR, i, tempResult, matrix)
		tempResult = deepSearch(iC, iR-1, i, tempResult, matrix)
		tempResult = deepSearch(iC, iR+1, i, tempResult, matrix)
	}

	return tempResult
}
