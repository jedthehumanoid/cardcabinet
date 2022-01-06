package cardcabinet

import (
	_"strings"
	"testing"
)

func TestMatch(t *testing.T) {
	tt := []struct {
		query string
		card string
		match bool
	}{
		{`labels "todo" ... labels "jobb" ... |`, "add-magnets.md", true},
		{`labels "todo"... labels "test" ... &`, "add-magnets.md", true},
		{`assignee "Ture Sventton" =`, "add-magnets.md", true},
	}

	for _, tc := range tt {
		t.Run(tc.query, func(t *testing.T) {
			card, err := ReadCard("testdata/"+tc.card)
			if err != nil {
				panic(err)
			}
			result := card.Match(tc.query)
			if result != tc.match {
				t.Errorf("unexpected result, %t != %t", result, tc.match)
			}
		
		
		})
	}
}
/*
func TestSplit(t *testing.T) {

	tt := []struct {
		query string
		ret   []string
	}{
		{"one two three", []string{"one", "two", "three"}},
		{"one \"forty two\" three", []string{"one", "forty two", "three"}},
		{"one 'forty two' three", []string{"one", "forty two", "three"}},
	}
	for _, tc := range tt {
		t.Run(tc.query, func(t *testing.T) {
			result := split(tc.query)
			if toJSON(result) != toJSON(tc.ret) {
				t.Fatalf("unexpected result, expected: %s, got: %s", toJSON(tc.ret), toJSON(result))
			}
		})
	}
}

func TestQueryRPN(t *testing.T) {
	tt := []struct {
		query []string
	}{
		{[]string{"labels", "todo", "...", "labels", "wontfix", "...", "&&"}},
		{[]string{"labels", "forty two", "...", "labels", "test", "...", "||"}},
		{[]string{"assignee", "Ture Sventton", "="}},
	}
	cards := ReadCards("testdata/", false)
	for _, tc := range tt {
		t.Run(strings.Join(tc.query, ","), func(t *testing.T) {
			queryRPN(cards[0], tc.query)
			/*
				if toJSON(result) != toJSON(tc.ret) {
					t.Fatalf("unexpected result, expected: %s, got: %s", toJSON(tc.ret), toJSON(result))
				}
			*/
			/*
		})
	}
}
*/
