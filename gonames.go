package gonames

import (
	// "fmt"
	"regexp"
	"strings"
)

var (

	// The order of RegEx matters!  Make sure the rules of
	// one doesn't invalidate the other...

	nameRegex = []map[string]func([][]string) string{

		{
			// Basic word boundary -- i.e. Dashes
			`[\w\']+`: func(matches [][]string) (word string) {

				// fmt.Printf("Dashes: %#v", matches)

				// Bypass the dashes
				for _, match := range matches {
					word += ucFirst(match[0]) + "-"
				}

				// Remove the final dash
				return strings.TrimRight(word, "-")
			},
		},
		{
			// Mc* (e.g. McDonald, McDowell)
			`^(.*)?(Mc)(.+)`: func(matches [][]string) (word string) {

				// fmt.Printf("Mc: %#v", matches)

				for i := 1; i < len(matches[0]); i++ {
					word += ucFirst(matches[0][i])
				}

				return strings.TrimSpace(word)
			},
		},
		{

			// O'* (e.g. O'Brian, O'Toole)
			`(?U)([^\']*)(O\')(\w+)\b+`: func(matches [][]string) (word string) {

				// fmt.Printf("O': %#v", matches)

				for _, match := range matches {
					for i := 1; i < len(match); i++ {
						word += ucFirst(match[i])
					}
				}

				return strings.TrimSpace(word)
			},
		},
	}
)

type NameMap map[string]string

func (n NameMap) GetFirstName() string {
	if _, ok := n["firstName"]; ok {
		return n["firstName"]
	}
	return ""
}

func (n NameMap) GetLastName() string {
	if _, ok := n["lastName"]; ok {
		return n["lastName"]
	}
	return ""
}

func (n NameMap) GetName() string {
	return strings.TrimSpace(n.GetFirstName() + " " + n.GetLastName())
}

func removeSpaces(name string) (result string) {

	for _, word := range strings.Split(name, " ") {
		if word != "" && word != " " {
			result += " " + word
		}
	}

	return strings.TrimSpace(result)
}

func createMap(in string) *NameMap {

	newMap := make(NameMap)

	names := strings.Split(in, " ")

	switch len(names) {
	case 1:
		newMap["firstName"] = strings.Join(names, " ")

	default:
		newMap["firstName"] = strings.Join(names[:len(names)-1], " ")
		newMap["lastName"] = strings.Join(names[len(names)-1:], "")
	}

	return &newMap
}

func ucFirst(word string) (out string) {

	splitWord := strings.Split(word, "")

	for _, letter := range word {
		return strings.ToUpper(string(letter)) + strings.Join(splitWord[len(string(letter)):], "")
	}

	return
}

func doFormat(in string) (out string) {

	for _, word := range regexp.MustCompile(`[^\w-\']`).Split(in, -1) {

		// Always capital first letter, regardless
		word = ucFirst(word)

		for _, rule := range nameRegex {
			for pattern, matchFunc := range rule {
				r := regexp.MustCompile(pattern)
				if m := r.FindAllStringSubmatch(word, -1); len(m) > 0 {
					word = matchFunc(m)
				}
			}
		}

		// Add the word to the output, plus a space
		out += word + " "

	}

	// Remove the final spaces
	return strings.TrimSpace(out)

}

func formatNames(names *NameMap) *NameMap {

	deref := *names

	if names.GetFirstName() != "" {
		deref["firstName"] = doFormat(deref["firstName"])
	}

	if names.GetLastName() != "" {
		deref["lastName"] = doFormat(deref["lastName"])
	}

	return names

}

func New(name string) (names *NameMap) {

	names = createMap(removeSpaces(name))
	formatNames(names)

	return names
}
