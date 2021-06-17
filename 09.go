package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("aaabbbbb\nkkkkkk")
	fmt.Println("aaabbbbb\\nkkkkkk")
	fmt.Println(`aaabbbbb\nkkkkkk`)
	fmt.Println(`aaabbbbb
						\nkkkkkk`)
	a := '*'
	fmt.Println(a)
	b := "abc"
	b = "cccc"
	fmt.Println(b)

	message := "shalom"
	for _, i := range message {
		fmt.Printf("%c\n", i)
	}

	message1 := "uv vagreangvbany fcnpr fgngvba"
	for i, i2 := range message1 {
		if i2 >= 'a' && i2 <= 'z' {
			i2 += 13
			if i2 > 'z' {
				i2 -= 26
			}
		}
		fmt.Printf("i is %v %c\n", i, i2)
	}

	question := "山东菏泽曹县"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)
}
