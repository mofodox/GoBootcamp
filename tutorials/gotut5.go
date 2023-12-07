package tutorials

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

/*
	Tutorial five covers on Strings and partial Rune types
*/

func TutorialFive() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)

	log.Println(sV2)

	log.Println("Length:", len(sV2))
	log.Println("Contains Another:", strings.Contains(sV2, "Another"))
	log.Println("o index:", strings.Index(sV2, "o"))
	log.Println("Replace:", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)
	log.Println(sV3)
	log.Println("Split:", strings.Split("a-b-c-d", "-"))
	log.Println("Lower:", strings.ToLower(sV2))
	log.Println("Upper:", strings.ToUpper(sV2))
	log.Println("Prefix:", strings.HasPrefix("tacocat", "taco"))
	log.Println("Suffix:", strings.HasSuffix("tacocat", "cat"))

	// Runes
	rStr := "abcdefg"
	log.Println("Rune count:", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d: %#U : %c\n", i, runeVal, runeVal)
	}
}