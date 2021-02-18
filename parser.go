package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parse(path string, alertions *Alertions) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == ".Unicode.empty()." {
			continue
		}
		seporator := "error:"
		sepPos := strings.Index(scanner.Text(), seporator)

		if sepPos <= -1 {
			seporator = "warning:"
			sepPos = strings.Index(scanner.Text(), seporator)
		}

		if sepPos > -1 {
			currentAllertionFileName := strings.Split(scanner.Text(), ":")[0]
			alertionPos := alertions.findIndexByName(currentAllertionFileName)
			if alertionPos != -1 {
				(*alertions)[alertionPos].text = append((*alertions)[alertionPos].text, strings.Split(scanner.Text(), seporator)[1])
				(*alertions)[alertionPos].strNumbers = append((*alertions)[alertionPos].strNumbers, strings.Split(scanner.Text(), ":")[1])
			} else {
				*alertions = append(*alertions, Alertion{
					fileName:   currentAllertionFileName,
					text:       []string{strings.Split(scanner.Text(), seporator)[1]},
					strNumbers: []string{strings.Split(scanner.Text(), ":")[1]},
				})
			}
		} else {
			panic(fmt.Sprintf("cannot determine is error or warning"))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
