package stringsample

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func PrintBytes(s string) {
	testWithinPackage()
	fmt.Println("........................")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\n")
	fmt.Println("=============================")
}

func PrintChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Printf("\n")
	fmt.Println("=============================")
}

func PrintChars1(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Printf("\n")
	fmt.Println("=============================")
}

func PrintString(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
	fmt.Println("=============================")

	byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)
}

func Length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func Mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func Split(str, sep string) []string {
	var result = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		result = append(result, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	result = append(result, str)
	return result
}

func Fib(n int) int {
	if n < 2 {
		return n
	} else {
		return Fib(n-1) + Fib(n-2)
	}

}
