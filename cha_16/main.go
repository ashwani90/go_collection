package main


import (
	"fmt"
	"strings"
	"regexp"
	"unicode"
)

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

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

	// Trimming character Sets
	trimmed = strings.Trim(description, "Asno ")
	fmt.Println("Trimmed:", trimmed)

	prefixedTrimmed := strings.TrimPrefix(description, "A boat ")
	wrongPrefix := strings.TrimPrefix(description, "A hat ")
	fmt.Println("Trimmed:", prefixedTrimmed)
	fmt.Println("Not trimmed:", wrongPrefix)

	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}

	trimmed = strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed:", trimmed)

	text := "It was a boat. A small boat."
	replace := strings.Replace(text, "boat", "canoe", 1)
	replaceAll := strings.ReplaceAll(text, "boat", "truck")
	fmt.Println("Replace:", replace)
	fmt.Println("Replace All:", replaceAll)

	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}
		return r
	}

	mapped := strings.Map(mapper, text)
	fmt.Println("Mapped:", mapped)

	// Using a string replacer
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	replaced := replacer.Replace(text)
	fmt.Println("Replaced: ", replaced)

	elements := strings.Fields(text)
	joined := strings.Join(elements, "--")
	fmt.Println("Joined:", joined)

	var builder strings.Builder
	for _, sub := range strings.Fields(text) {
		if (sub == "small") {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}
	fmt.Println("String: ", builder.String())

	match, err := regexp.MatchString("[A-z]oat", description)
	if (err == nil) {
		fmt.Println("Match: ", match)
	} else {
		fmt.Println("Error: ", err)
	}

	pattern, compileErr := regexp.Compile("[A-z]oat")
	description = "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"

	if (compileErr == nil) {
		fmt.Println("Description: ", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error: ", compileErr)
	}

	pattern = regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description = "Kayak. A boat for one person"
	firstIndex := pattern.FindStringIndex(description)
	allIndices := pattern.FindAllStringIndex(description, -1)
	fmt.Println("First index", firstIndex[0], "-", firstIndex[1], "=", getSubstring(description, firstIndex))
	
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-", idx[1], "=", getSubstring(description, idx))
	}

	firstMatch := pattern.FindString(description)
	allMatches := pattern.FindAllString(description, -1)
	fmt.Println("First match:", firstMatch)

	for i, m := range allMatches {
		fmt.Println("Match", i, "=", m)
	}

	pattern = regexp.MustCompile(" |boat|one")
	split := pattern.Split(description, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}

	pattern = regexp.MustCompile("A [A-z]* for [A-z]* person")
	description = "Kayak. A boat for one person"
	str := pattern.FindString(description)
	fmt.Println("Match:", str)

	subs := pattern.FindStringSubmatch(description)
	for _, s:= range subs {
		fmt.Println("Match:", s)
	}

	pattern = regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	subs = pattern.FindStringSubmatch(description)
	for _, name := range []string {"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}

	template := "(type: ${type}, capacity ${capacity})"
	replaced = pattern.ReplaceAllString(description, template)
	fmt.Println(replaced)

	replaced = pattern.ReplaceAllStringFunc(description, func(s string) string {
		return "This is the replacement content"
	})
	fmt.Println(replaced)
}