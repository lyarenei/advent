package _5

import (
	"2024/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type PrintRules map[int][]int

type PrintOrder struct {
	PageOrder []int
}

func (po PrintOrder) GetMiddleValue() int {
	// Don't know rules for even lengths, leaving as is
	idx := len(po.PageOrder) / 2
	return po.PageOrder[idx]
}

func Run(inputFile string) {
	rules, orders := readFile(inputFile)
	sum := 0
	for _, order := range orders {
		ok := isValid(rules, order.PageOrder)
		if ok {
			sum += order.GetMiddleValue()
		}
	}

	fmt.Printf("The sum of all middle page numbers in valid printing rules is %d\n", sum)
}

func readFile(fileName string) (PrintRules, []PrintOrder) {
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

	rules := make(PrintRules)
	var orders []PrintOrder
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			page := utils.StringToInt(parts[0])
			before := utils.StringToInt(parts[1])
			rules[page] = append(rules[page], before)
			continue
		}

		pages := strings.Split(line, ",")
		orders = append(orders, PrintOrder{
			PageOrder: utils.MapSlice(pages, utils.StringToInt),
		})
	}

	return rules, orders
}

func isValid(rules PrintRules, order []int) bool {
	if len(rules) == 0 || len(order) <= 1 {
		return true
	}

	firstPage := order[0]
	nextPage := order[1]
	pageRules, ok := rules[nextPage]
	if ok {
		if slices.Contains(pageRules, firstPage) {
			return false
		}
	}

	return isValid(rules, order[1:])
}
