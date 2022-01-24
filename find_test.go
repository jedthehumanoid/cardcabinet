package cardcabinet

import "testing"

func TestFind(t *testing.T) {
	expected := []string{
		"testdata/add-magnets.md",
		"testdata/card.json",
		"testdata/card.md",
		"testdata/types.md",
		"testdata/without-frontmatter.json",
		"testdata/without-frontmatter.md",
		"testdata/wonky-middlematter.json",
		"testdata/wonky-middlematter.md",
	}
	result := findFiles("testdata/", false)

	if toJSON(expected) != toJSON(result) {
		t.Errorf("unexpected result, expected:\n%s\ngot:\n%s\n", expected, result)
	}
}

func TestFindRecursive(t *testing.T) {
	expected := []string{
		"testdata/add-magnets.md",
		"testdata/card.json",
		"testdata/card.md",
		"testdata/recursive/recursive.md",
		"testdata/types.md",
		"testdata/without-frontmatter.json",
		"testdata/without-frontmatter.md",
		"testdata/wonky-middlematter.json",
		"testdata/wonky-middlematter.md",
	}
	result := findFiles("testdata/", true)

	if toJSON(expected) != toJSON(result) {
		t.Errorf("unexpected result, expected:\n%s\ngot:\n%s\n", expected, result)
	}
}
