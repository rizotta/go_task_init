package main

import (
	"fmt"
	"strings"
)

func main() {
	var strToCheck string
	fmt.Scan(&strToCheck)
	fmt.Println(checkBrackets(strToCheck))
}

func checkBrackets(str string) bool {
	var brStack []string
	brMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	slStr := strings.Split(str, "")
	for _, v := range slStr {
		if v == "(" || v == "{" || v == "[" {
			brStack = append(brStack, v)
		} else if v == ")" || v == "}" || v == "]" {
			if len(brStack) == 0 || brStack[len(brStack)-1] != brMap[v] {
				return false
			} else {
				brStack = brStack[:len(brStack)-1]
			}
		}
	}

	if len(brStack) > 0 {
		return false
	}

	return true
}
