package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	str := "Yūgen ☯ 💀"

	bytes := []byte(str)
	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n\n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n\n", utf8.RuneCount(bytes))

	for i := range str {
		fmt.Printf("str[%2d] = %-2x\n", i, str[i])
	}
	fmt.Println()
	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	// slice bytes
	fmt.Printf("\nGiven string: %s\n", str)
	fmt.Printf("1st Byte %c\n", str[0])
	fmt.Printf("2nd Byte %c\n\n", str[1]) // notice second value is not same in string

	fmt.Printf("2nd Rune %s\n", str[1:3])            // second rune, prints second char of string
	fmt.Printf("last Rune %s\n\n", str[len(str)-4:]) // last rune which has 4 bytes

	// rune slice to avoid char breaks

	runes := []rune(str)
	fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d runes\n\n", len(runes))

	//indexing and slicing runse
	fmt.Printf("1st Rune %c\n", runes[0])
	fmt.Printf("2nd Rune %c\n\n", runes[1])

	// rune slice first five chars
	fmt.Printf("First five chars %c\n", runes[:5])

	// avoid rune slices bcoz its costly op
	fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d bytes (size of runes slice)\n",
		int(unsafe.Sizeof(runes[0]))*len(runes))
	fmt.Printf("\t%d Length of runes\n", len(runes))

	// But notice size of bytes less
	fmt.Printf("\t%d bytes (size of bytes slice)\n",
		int(unsafe.Sizeof(bytes[0]))*len(bytes))

	// use for range for string loops, it will auto converts to rune char
}

//  %x (hexadecimal) format
//  % space x to disply bytes with spaces between
// %q (quoted) verb will escape any non-printable byte sequences in a string
