package _1_basics

import "fmt"

func Strings() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	//uint8 == byte
	// strings in go are immutable

	fmt.Println(book[4:11])

	//slice (no end)
	fmt.Println(book[4:])

	//slice (no start)
	fmt.Println(book[:4])

	// `+` for concatenation
	fmt.Println("t" + book[1:])

	// multi line

	poem := `
		The road goes ever on
		Down from the door where it began
		...
	`
	fmt.Println(poem)
}
