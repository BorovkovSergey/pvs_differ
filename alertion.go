package main

// import(
//     "fmt"
// )

type Alertions []Alertion
type Strings []string

type Alertion struct {
	fileName   string
	text       Strings
	strNumbers Strings
	isError    bool
}

func (e Alertions) findIndexByName(name string) int {
	for i, alert := range e {
		if alert.fileName == name {
			return i
		}
	}
	return -1
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func (e Strings) Contains(x string) bool {
	for _, n := range e {
		if x == n {
			return true
		}
	}
	return false
}

func (e *Alertions) Contains(input Alertion) bool {
	pos := (*e).findIndexByName(input.fileName)
	if pos != -1 {
		return true
	}
	return false
}

func (e *Alertions) inputIfNotContain(input Alertion) bool {
	if !e.Contains(input) {
		*e = append(*e, input)
		return true
	}
	return false
}

func (e *Alertion) inputTextIfNotContain(input string) {
	if !(*e).text.Contains(input) {
		(*e).text = append((*e).text, input)
	}
}

func (e Strings) makeDiff(inputs Strings, positions *[]int) Strings {
	diff := Strings{}
	for i, input := range inputs {
		if !e.Contains(input) {
			diff = append(diff, input)
			*positions = append(*positions, i)
		}
	}
	return diff
}

func (e *Strings) addSlice(inputs Strings) {
	for _, i := range inputs {
		(*e) = append((*e), i)
	}
}

func MakeStrNumbersDiff(input Strings, positions []int) Strings {
	res := Strings{}
	for _, pos := range positions {
		res = append(res, input[pos])
	}
	return res
}
func (e Alertions) MakeAlertionsDiff(inputs Alertions) Alertions {
	diff := Alertions{}
	for _, input := range inputs {
		pos := e.findIndexByName(input.fileName)
		if pos != -1 {
			strNumbersPos := []int{}
			diffStr := e[pos].text.makeDiff(input.text, &strNumbersPos)
			diffStrNumbers := MakeStrNumbersDiff(input.strNumbers, strNumbersPos)
			pos = diff.findIndexByName(input.fileName)
			if len(diffStr) <= 0 {
				continue
			}
			if pos != -1 {
				diff[pos].text.addSlice(diffStr)
				diff[pos].strNumbers.addSlice(diffStrNumbers)
			} else {
				diff = append(diff, Alertion{
					fileName:   input.fileName,
					text:       diffStr,
					strNumbers: diffStrNumbers,
				})
			}
		} else {
			diff = append(diff, input)
		}
	}
	return diff
}
