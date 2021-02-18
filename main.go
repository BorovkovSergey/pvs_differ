package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("Usage: %s <original file> <new file>\n", args[0])
		return
	}
	originalFile := os.Args[1]
	newFile := os.Args[2]

	oldAlertions := Alertions{}
	parse(originalFile, &oldAlertions)
	newAlertions := Alertions{}
	parse(newFile, &newAlertions)

	diffAlertions := oldAlertions.MakeAlertionsDiff(newAlertions)
	count := 0
	for i, v := range diffAlertions {
		fmt.Printf("%d: %s\n", i, v.fileName)
		for j, s := range v.text {
			count++
			fmt.Printf("\t\t%d: string[%s] %s\n", count, v.strNumbers[j], s)
		}
	}
}
