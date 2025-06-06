package utils

import (
	"strconv"
	"strings"
	"unicode"
)

import "regexp"

var validExpr = regexp.MustCompile(`^[\d\s+\-*/()]+$`)

func Calculate(expr string) int {
	if strings.TrimSpace(expr) == "" || !validExpr.MatchString(expr) {
		return 0
	}
	var numStack []int
	var opStack []rune
	priority := map[rune]int{
		'+': 1, '-': 1,
		'*': 2, '/': 2,
	}

	eval := func() {
		b := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		a := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]

		switch op {
		case '+':
			numStack = append(numStack, a+b)
		case '-':
			numStack = append(numStack, a-b)
		case '*':
			numStack = append(numStack, a*b)
		case '/':
			numStack = append(numStack, a/b)
		}
	}

	for i := 0; i < len(expr); i++ {
		ch := rune(expr[i])
		if ch == ' ' {
			continue
		}
		if unicode.IsDigit(ch) {
			n := 0
			for i < len(expr) && unicode.IsDigit(rune(expr[i])) {
				digit, _ := strconv.Atoi(string(expr[i]))
				n = n*10 + digit
				i++
			}
			i--
			numStack = append(numStack, n)
		} else if ch == '(' {
			opStack = append(opStack, ch)
		} else if ch == ')' {
			for opStack[len(opStack)-1] != '(' {
				eval()
			}
			opStack = opStack[:len(opStack)-1]
		} else {
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' &&
				priority[opStack[len(opStack)-1]] >= priority[ch] {
				eval()
			}
			opStack = append(opStack, ch)
		}
	}

	for len(opStack) > 0 {
		eval()
	}

	return numStack[len(numStack)-1]
}
