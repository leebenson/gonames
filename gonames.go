package gonames

import (
	// "fmt"
	"regexp"
	"strings"
)

type NameMap struct {
	firstName string
	lastName  string
}

func (n *NameMap) GetFirstName() string {
	return n.firstName
}

func (n *NameMap) GetLastName() string {
	return n.lastName
}

func (n *NameMap) GetName() string {
	return strings.TrimSpace(n.firstName + " " + n.lastName)
}

func (n *NameMap) SetFirstName(name string) *NameMap {
	n.firstName = doFormat(name)
	return n
}

func (n *NameMap) SetLastName(name string) *NameMap {
	n.lastName = doFormat(name)
	return n
}

func removeSpaces(name string) (result string) {

	for _, word := range strings.Split(name, " ") {
		if word != "" && word != " " {
			result += " " + word
		}
	}

	return strings.TrimSpace(result)
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

func New() (names *NameMap) {
	return new(NameMap)
}
