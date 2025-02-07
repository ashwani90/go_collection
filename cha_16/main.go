package main


import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	product := "Kayak"
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("Contains Any:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))

	description := "A boat is sailing"
	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.Title(description))

	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))

	// Working with Characters
	for _, char := range product {
		fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
	}

	description2 := "A boat for one person"
	fmt.Println("Count:", strings.Count(description2, "o"))
	fmt.Println("Index:", strings.Index(description2, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description2, "o"))
	fmt.Println("IndexAny:", strings.IndexAny(description2, "abcd"))
	fmt.Println("LastIndex:", strings.LastIndex(description2, "o"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description2, "abcd"))

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))

	splits := strings.Split(description2, " ")
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}

	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}
	description = "This  is  double  spaced"
	splits2 := strings.SplitN(description, " ", 3)
	for _, x := range splits2 {
		fmt.Println("SplitAfter >> " + x + "<<")
	}

	splitter := func(r rune) bool {
		return r == ' '
	}

	splits = strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}

	username := "Alice"
	trimmed := strings.TrimSpace(username)
	fmt.Println("Trimmed:", ">>" + trimmed + "<<")
}