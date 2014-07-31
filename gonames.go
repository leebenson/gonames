package gonames

import (
	// "fmt"
	"regexp"
	"strings"
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

	r := regexp.MustCompile(`[\w\']+-?`)

	for _, matches := range r.FindAllStringSubmatch(in, -1) {
		for _, word := range matches {

			// Uppercase the first letter, regardless
			word = ucFirst(word)

			// Mc*, O'*
			for _, pattern := range []string{`^(Mc)(.+)$`, `^(O\')(.+)$`} {
				r := regexp.MustCompile(pattern)
				if m := r.FindAllStringSubmatch(word, -1); m != nil {

					word = func(matches [][]string) (word string) {

						// fmt.Printf("O': %#v", matches)

						for _, match := range matches {
							for i := 1; i < len(match); i++ {
								word += ucFirst(match[i])
							}
						}

						return strings.TrimSpace(word)
					}(m)
				}
			}

			out += word + " "
		}
	}

	// Glue back any dashes, and trim the fat
	return strings.TrimSpace(strings.Replace(out, "- ", "-", -1))
}

func formatNames(names NameMap) *NameMap {

	if names.GetFirstName() != "" {
		names["firstName"] = doFormat(names["firstName"])
	}

	if names.GetLastName() != "" {
		names["lastName"] = doFormat(names["lastName"])
	}

	return &names

}

func New(name string) (names *NameMap) {

	names = createMap(removeSpaces(name))
	formatNames(*names)

	return names
}
