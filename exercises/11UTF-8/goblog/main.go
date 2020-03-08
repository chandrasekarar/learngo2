package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Byte loop: with %q")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Byte loop: with %+q")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%+q", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	strBytes := []byte(sample)
	fmt.Println("byte slice Println:")
	fmt.Println(strBytes)
	fmt.Println("byte slice Printf with % x:")
	fmt.Printf("% x\n", strBytes)

	fmt.Println("byte slice Printf with %q:")
	fmt.Printf("%q\n", strBytes)

	const nihongo = "日本語\x3d"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
