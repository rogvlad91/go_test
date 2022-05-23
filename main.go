package main

import (
	"bufio"
	"fmt"
	"os"
)

func getReversedHook(bracket rune) (reversedBracket rune) {
	switch bracket {
	case '}':
		return '{'
	case ')':
		return '('
	case ']':
		return '['
	default:
		break
	}
	return
}

func isValidParentheses(brackets string) (flag bool) {
	stack := make([]rune, 0)

	for _, hook := range brackets {
		switch hook {
		case '(', '{', '[':
			stack = append(stack, hook)
		case ')', '}', ']':
			reversedHook := getReversedHook(hook)
			if len(stack) == 0 || stack[len(stack)-1] != reversedHook {
				flag = false
				return
			}
			stack = stack[:len(stack)-1]
		default:
			break
		}
	}

	if len(stack) == 0 {
		flag = true
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter bracket string: ")
	brackets, _ := reader.ReadString('\n')
	flag := isValidParentheses(brackets)

	fmt.Println(flag)
}
