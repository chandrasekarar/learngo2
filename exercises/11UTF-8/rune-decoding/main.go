package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Printf("\n%33s\n%s\n\n", "Rune Decoding", strings.Repeat("-", 55))

	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	fmt.Println(text)
	fmt.Println()

	// manually decode the runes
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)

		i += size
	}
	fmt.Printf("\n\n")

	// automatically decode rune using for range

	for _, r := range text {
		fmt.Printf("%c", r)
	}
}
