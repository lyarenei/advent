package _5

import (
	"2024/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PrintRules map[int][]int

type PrintOrder struct {
	PageOrder []int
}

func Run(inputFile string) {
	rules, orders := readFile(inputFile)
	fmt.Printf("Read printing rules for %d pages, there are %d printing orders\n", len(rules), len(orders))
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
