package gonames

import "testing"

var nameSet = []map[string]string{

	// BASIC NAMES

	// 1 first name
	{
		"first":         "jamie",
		"last":          "",
		"expectedFirst": "Jamie",
		"expectedLast":  "",
		"expectedFull":  "Jamie",
	},

	// 1 first name, 1 last name
	{
		"first":         "james",
		"last":          "jones",
		"expectedFirst": "James",
		"expectedLast":  "Jones",
		"expectedFull":  "James Jones",
	},

	// 2 first, 1 last name
	{
		"first":         "susan louise ",
		"last":          "logan",
		"expectedFirst": "Susan Louise",
		"expectedLast":  "Logan",
		"expectedFull":  "Susan Louise Logan",
	},

	// 3 first, 1 last name
	{
		"first":         "matthew andrew lloyd",
		"last":          "davies",
		"expectedFirst": "Matthew Andrew Lloyd",
		"expectedLast":  "Davies",
		"expectedFull":  "Matthew Andrew Lloyd Davies",
	},

	// DASHES
	{
		"first":         "  john ",
		"last":          "david-smith   ",
		"expectedFirst": "John",
		"expectedLast":  "David-Smith",
		"expectedFull":  "John David-Smith",
	},

	// Mc*
	{
		"first":         "jonathan",
		"last":          "mcdonald   ",
		"expectedFirst": "Jonathan",
		"expectedLast":  "McDonald",
		"expectedFull":  "Jonathan McDonald",
	},

	// Mc*, dashes, 2 first names

	{
		"first":         "patricia june",
		"last":          "henry-mcdonald ",
		"expectedFirst": "Patricia June",
		"expectedLast":  "Henry-McDonald",
		"expectedFull":  "Patricia June Henry-McDonald",
	},

	// O'*
	{
		"first":         "     daniel    ",
		"last":          "o'toole ",
		"expectedFirst": "Daniel",
		"expectedLast":  "O'Toole",
		"expectedFull":  "Daniel O'Toole",
	},

	{
		"first":         "     rebecca    ",
		"last":          "  smith-mcdonald-o'toole-o'leary",
		"expectedFirst": "Rebecca",
		"expectedLast":  "Smith-McDonald-O'Toole-O'Leary",
		"expectedFull":  "Rebecca Smith-McDonald-O'Toole-O'Leary",
	},

	// O'*-*

	{
		"first":         "     lee    ",
		"last":          "o'brian-keith",
		"expectedFirst": "Lee",
		"expectedLast":  "O'Brian-Keith",
		"expectedFull":  "Lee O'Brian-Keith",
	},

	// O'*-Mc*
	{
		"first":         "theodore",
		"last":          "o'brian-mcdonald",
		"expectedFirst": "Theodore",
		"expectedLast":  "O'Brian-McDonald",
		"expectedFull":  "Theodore O'Brian-McDonald",
	},

	// Prefixes
	{
		"first":         "james",
		"last":          "van der beek",
		"expectedFirst": "James",
		"expectedLast":  "Van Der Beek",
		"expectedFull":  "James Van Der Beek",
	},
}

func TestNames(t *testing.T) {

	// Only need to run each rule once, since the RegEx is in
	// slice order...

	for _, name := range nameSet {

		result := New()
		result.SetFirstName(name["first"])
		result.SetLastName(name["last"])

		// t.Log("Testing " + name + "... (" + result.GetName() + ")")

		// First name
		if result.GetFirstName() != name["expectedFirst"] {
			t.Errorf(`Expected: "%s", got: "%s"`, name["expectedFirst"], result.GetFirstName())
		}

		// Last name
		if result.GetLastName() != name["expectedLast"] {
			t.Errorf(`Expected: "%s", got: "%s"`, name["expectedLast"], result.GetLastName())
		}

		// Full name
		if result.GetName() != name["expectedFull"] {
			t.Errorf(`Expected: "%s", got: "%s"`, name["expectedFull"], result.GetName())
		}
	}

}

func Benchmark1Name(b *testing.B) {

	for i := 0; i < b.N; i++ {
		r := New()
		r.SetFirstName("james")
	}

}

func Benchmark2Names(b *testing.B) {

	for i := 0; i < b.N; i++ {
		r := New()
		r.SetFirstName("albert")
		r.SetLastName("einstein")
	}

}

func BenchmarkLongName1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		r := New()
		r.SetFirstName("kal-el")
		r.SetLastName("(son of) jor-el")
	}

}

func BenchmarkLongName2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		r := New()
		r.SetFirstName("old mcdonald had a farm")
		r.SetLastName("but he wasn't called mcdonald-o'brian")
	}

}
