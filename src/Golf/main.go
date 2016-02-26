package main

import (
	"Golf/golf"
	"bufio"

	"os"
)

func main() {
	file_input, _ := os.OpenFile("inputGolf.txt", os.O_RDONLY, 0755)

	scanner := bufio.NewScanner(file_input)
	table := golf.NewTable(8, 5, scanner)

	file_input.Close()

	table.Solve()
	table.Print()
	table.PrintResult()
}
