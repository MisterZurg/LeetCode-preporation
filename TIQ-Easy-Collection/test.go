package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("ABOBA\n\t")
	fmt.Println(`ABOBA\n\t`)

	emojo := "🙋🌍❗"
	fmt.Println(emojo)
	fmt.Println(len(emojo))
	fmt.Println(utf8.RuneCountInString(emojo))

}
