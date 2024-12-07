package _6

import (
	"2024/utils"
	"bufio"
	"fmt"
	"os"
)

func Run(inputFile string) {
	labMap := readFile(inputFile)
	predictRoute(labMap)
	positions := utils.Search2DSimple(labMap, 'x')
	fmt.Printf("The guard visited %d unique positions before leaving the lab\n", positions)
}

func printMap(arr [][]rune) {
	for _, row := range arr {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func readFile(fileName string) [][]rune {
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var arr [][]rune
	for scanner.Scan() {
		arr = append(arr, []rune(scanner.Text()))
	}

	return arr
}

type Movement int

const (
	Up Movement = iota
	Down
	Left
	Right
)

func predictRoute(arr [][]rune) {
	currentDirection := Up
	guardAppearance := '^'
	for {
		// Find guard position
		row, col := utils.FindIdx2D(arr, guardAppearance)
		if row == -1 && col == -1 {
			printMap(arr)
			return
		}

		guardMove(arr, currentDirection, row, col)
		switch currentDirection {
		case Up:
			currentDirection = Right
			guardAppearance = '>'
		case Right:
			currentDirection = Down
			guardAppearance = 'v'
		case Down:
			currentDirection = Left
			guardAppearance = '<'
		case Left:
			currentDirection = Up
			guardAppearance = '^'
		}
	}
}

func guardMove(arr [][]rune, direction Movement, row, col int, guard rune) {
	var line []rune
	var axis utils.Direction
	switch direction {
	case Up:
		axis = utils.Vertical
		// Get line from top edge up to guard
		line = utils.GetLine(arr, axis, 0, col, row+1)
		// Find first obstacle for guard (guard is at the last position in line)
		obstacleIdx := utils.LastIndex(line, '#')
		// Move the guard to the obstacle
		moveGuard(line, direction, obstacleIdx+1, len(line)-1)
		// Update map
		utils.SetLine2D(arr, axis, 0, col, line)
	case Right:
		axis = utils.Horizontal
		// Get line from guard up to the right edge
		line = utils.GetLine(arr, axis, row, col, len(arr[row]))
		// Find first obstacle for guard (guard is at the first position in line)
		obstacleIdx := utils.FirstIndex(line, '#')
		// Move the guard to the obstacle
		moveGuard(line, direction, 0, obstacleIdx-1)
		// Update map
		utils.SetLine2D(arr, axis, row, col, line)
	case Down:
		axis = utils.Vertical
		// Get line from guard down to the bottom edge
		line = utils.GetLine(arr, axis, row, col, len(arr))
		// Find first obstacle for guard (guard is at the first position in line)
		obstacleIdx := utils.FirstIndex(line, '#')
		// Move the guard to the obstacle
		moveGuard(line, direction, 0, obstacleIdx-1)
		// Update map
		utils.SetLine2D(arr, axis, row, col, line)
	case Left:
		axis = utils.Horizontal
		// Get line from left edge up to the guard
		line = utils.GetLine(arr, axis, row, 0, col+1)
		// Find first obstacle for guard (guard is at the last position in line)
		obstacleIdx := utils.LastIndex(line, '#')
		// Move the guard (the obstacle is at obstacleIdx)
		moveGuard(line, direction, obstacleIdx+1, len(line)-1)
		// Update map
		utils.SetLine2D(arr, axis, row, 0, line)
	}
}

func moveGuard(path []rune, direction Movement, startIdx int, endIdx int) {
	sIdx := startIdx
	if startIdx < 0 {
		sIdx = 0
	}

	eIdx := endIdx
	if endIdx < 0 {
		eIdx = len(path) - 1
	}

	for i := sIdx; i <= eIdx; i++ {
		path[i] = 'x'
	}

	if startIdx < 0 || endIdx < 0 {
		return
	}

	switch direction {
	case Up:
		path[startIdx] = '>'
	case Down:
		path[endIdx] = '<'
	case Left:
		path[startIdx] = '^'
	case Right:
		path[endIdx] = 'v'
	}
}
