package main

import "fmt"

func match(regexp, text string) bool {
	if len(regexp) > 1 && regexp[0] == '^' {
		return matchHere(regexp[1:], text)
	}
	if matchHere(regexp, text) {
		return true
	}
	for {
		if len(text) == 0 {
			break
		}
		if matchHere(regexp, text) {
			return true
		}
		text = text[1:]
	}
	return false
}

func matchHere(regexp, text string) bool {
	if regexp == "" {
		return true
	}
	if len(regexp) > 1 && (regexp[1] == '*') {
		return matchStar(regexp[0], regexp[2:], text)
	}
	if regexp == "$" {
		return text == ""
	}
	if text != "" && (regexp[0] == '.' || regexp[0] == text[0]) {
		return matchHere(regexp[1:], text[1:])
	}
	return false
}

func matchStar(char uint8, regexp, text string) bool {
	if matchHere(regexp, text) {
		return true
	}
	for {
		if len(text) == 0 {
			break
		}
		if text[0] == char || char == '.' {
			if matchHere(regexp, text) {
				return true
			}
		} else {
			return false
		}
		text = text[1:]
	}
	return false
}

func main()  {
	fmt.Println(match(".*", "a"))
	fmt.Println(match("a*b", "aab"))
	fmt.Println(match("a*c", "aab"))
	fmt.Println(match("^match end$", "match end"))
	fmt.Println(match("match end", "match end"))
	fmt.Println(match("a*c$", "aac"))
}