package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Hello Gophercies 😁", 40)

	x, y := 1, "1"
	fmt.Printf("x = %v, y = %v\n", x, y)
	fmt.Printf("x = %#v, y = %#v\n", x, y) // %#v prints the logging format, use to debug //actuall value

	fmt.Print(isPalindrome("radar"))
}

func isPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] { // len(s) is the length of the string
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2 BUG: this is in bytes, not runes
	padding := (width - utf8.RuneCountInString((text))) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
