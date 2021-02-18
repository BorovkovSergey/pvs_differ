package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parse(path string) Alertions {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := Alertions{}
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
			alertionPos := res.findIndexByName(currentAllertionFileName)
			if alertionPos != -1 {
				res[alertionPos].text = append(res[alertionPos].text, seporator+strings.Split(scanner.Text(), seporator)[1])
				res[alertionPos].strNumbers = append(res[alertionPos].strNumbers, strings.Split(scanner.Text(), ":")[1])
			} else {
				res = append(res, Alertion{
					fileName:   currentAllertionFileName,
					text:       []string{seporator + strings.Split(scanner.Text(), seporator)[1]},
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
	return res
}
