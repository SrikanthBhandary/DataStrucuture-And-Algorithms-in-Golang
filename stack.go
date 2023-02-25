package main

import (
	"errors"
	"fmt"
)

type Linter struct {
	stack        []rune
	BraceMapping map[rune]rune
}

func NewLinter() *Linter {
	m := make(map[rune]rune)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'
	return &Linter{BraceMapping: m}
}

func IsOpeningBrace(char rune) bool {
	switch char {
	case '(':
	case '[':
	case '{':
		return true
	default:
		return false

	}
	return true
}

func IsClosingBrace(char rune) bool {
	switch char {
	case ')':
	case ']':
	case '}':
		return true
	default:
		return false

	}
	return true
}

func (L *Linter) Lint(text string) error {

	for i, char := range text {
		if IsOpeningBrace(char) {
			L.stack = append(L.stack, char)
		} else if IsClosingBrace(char) {
			if len(L.stack) > 0 && L.stack[len(L.stack)-1] == L.BraceMapping[char] {
				L.stack = L.stack[:len(L.stack)-1]
			} else {
				return errors.New(fmt.Sprintf("Syntax not correct at char %d", i))
			}

		}
	}
	return nil

}

func main() {
	linter := NewLinter()
	err := linter.Lint("(())")
	fmt.Println(err)

}
