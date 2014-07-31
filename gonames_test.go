package gonames

import "testing"

var nameSet = map[string]map[string]string{

	// BASIC NAMES

	// 1 first, 0 last name
	"jamie": map[string]string{
		"expectedFirst": "Jamie",
		"expectedLast":  "",
		"expectedFull":  "Jamie",
	},

	// 1 first name, 1 last name
	"james jones": map[string]string{
		"expectedFirst": "James",
		"expectedLast":  "Jones",
		"expectedFull":  "James Jones",
	},

	// 2 first, 1 last name
	"susan louise logan": map[string]string{
		"expectedFirst": "Susan Louise",
		"expectedLast":  "Logan",
		"expectedFull":  "Susan Louise Logan",
	},

	// 3 first, 1 last name
	"matthew andrew lloyd davies": map[string]string{
		"expectedFirst": "Matthew Andrew Lloyd",
		"expectedLast":  "Davies",
		"expectedFull":  "Matthew Andrew Lloyd Davies",
	},

	// DASHES
	"  john david-smith  ": map[string]string{
		"expectedFirst": "John",
		"expectedLast":  "David-Smith",
		"expectedFull":  "John David-Smith",
	},

	// Mc*
	"jonathan mcdonald": map[string]string{
		"expectedFirst": "Jonathan",
		"expectedLast":  "McDonald",
		"expectedFull":  "Jonathan McDonald",
	},

	// Mc*, dashes, 2 first names
	"patricia june henry-mcdonald": map[string]string{
		"expectedFirst": "Patricia June",
		"expectedLast":  "Henry-McDonald",
		"expectedFull":  "Patricia June Henry-McDonald",
	},

	// O'*
	"daniel      o'toole": map[string]string{
		"expectedFirst": "Daniel",
		"expectedLast":  "O'Toole",
		"expectedFull":  "Daniel O'Toole",
	},

	"rebecca smith-mcdonald-o'toole-o'leary": map[string]string{
		"expectedFirst": "Rebecca",
		"expectedLast":  "Smith-McDonald-O'Toole-O'Leary",
		"expectedFull":  "Rebecca Smith-McDonald-O'Toole-O'Leary",
	},

	// O'*-*
	"lee o'brian-keith": map[string]string{
		"expectedFirst": "Lee",
		"expectedLast":  "O'Brian-Keith",
		"expectedFull":  "Lee O'Brian-Keith",
	},

	// O'*-Mc*
	"theodore o'brian-mcdonald": map[string]string{
		"expectedFirst": "Theodore",
		"expectedLast":  "O'Brian-McDonald",
		"expectedFull":  "Theodore O'Brian-McDonald",
	},
}

func TestNames(t *testing.T) {

	// Only need to run each rule once, since the RegEx is in
	// slice order...

	for name, options := range nameSet {
		result := New(name)

		// t.Log("Testing " + name + "... (" + result.GetName() + ")")

		// First name
		if result.GetFirstName() != options["expectedFirst"] {
			t.Errorf(`Expected: "%s", got: "%s"`, options["expectedFirst"], result.GetFirstName())
		}

		// Last name
		if result.GetLastName() != options["expectedLast"] {
			t.Errorf(`Expected: "%s", got: "%s"`, options["expectedLast"], result.GetLastName())
		}

		// Full name
		if result.GetName() != options["expectedFull"] {
			t.Errorf(`Expected: "%s", got: "%s"`, options["expectedFull"], result.GetName())
		}
	}

}

func Benchmark1Name(b *testing.B) {

	name := "james"

	for i := 0; i < b.N; i++ {
		New(name)
	}

}

func Benchmark2Names(b *testing.B) {

	name := "albert einstein"

	for i := 0; i < b.N; i++ {
		New(name)
	}

}

func BenchmarkLongName1(b *testing.B) {

	name := "kal-el (son of) jor-el"

	for i := 0; i < b.N; i++ {
		New(name)
	}

}

func BenchmarkLongName2(b *testing.B) {

	name := "old mcdonald had a farm, but he wasn't called mcdonald-o'brian"

	for i := 0; i < b.N; i++ {
		go New(name)
	}

}
