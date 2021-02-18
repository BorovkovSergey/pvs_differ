package main

import (
	"fmt"
	"os"
)

func PrintDiff(diff Alertions) {
	count := 0
	for i, v := range diff {
		fmt.Printf("%d: %s\n", i, v.fileName)
		for j, s := range v.text {
			count++
			fmt.Printf("\t\t%d: string[%s] %s\n", count, v.strNumbers[j], s)
		}
	}
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("Usage: %s <original file> <new file>\n", args[0])
		return
	}
	originalFile := os.Args[1]
	newFile := os.Args[2]

	oldAlertions := parse(originalFile)
	newAlertions := parse(newFile)

	diffAlertions := oldAlertions.MakeAlertionsDiff(newAlertions)
	PrintDiff(diffAlertions)
}
