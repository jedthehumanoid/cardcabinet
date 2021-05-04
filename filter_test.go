package cardcabinet

import (
	"testing"
"strings"
)

func TestQuery(t *testing.T) {

	tt := []struct {
		query string
		cards []string
	}{
		{"labels todo ... labels test ... && labels jobb ... ||", []string{"add-magnets.md", "types.md"}},
		{"labels todo ... labels test ... &&", []string{"add-magnets.md"}},
		//{"name types.md =", []string{"types.md"}},
		{"assignee \"Ture Sventton\" =", []string{"add-magnets.md"}},
	}
	cards := ReadCards("testdata/")
	for _, tc := range tt {
		t.Run(tc.query, func(t *testing.T) {
			r := queryCards(cards, tc.query)
			result := []string{}
			for _, card := range r {
				result = append(result, card.Name)
			}
			if toJSON(result) != toJSON(tc.cards) {
				t.Fatalf("unexpected result, expected: %s, got: %s", toJSON(tc.cards), toJSON(result))
			}
		})
	}
}


func TestSplit(t *testing.T) {

	tt := []struct {
		query string
		ret []string
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
	cards := ReadCards("testdata/")
	for _, tc := range tt {
		t.Run(strings.Join(tc.query, ","), func(t *testing.T) {
			queryRPN(cards[0],tc.query)
/*
			if toJSON(result) != toJSON(tc.ret) {
				t.Fatalf("unexpected result, expected: %s, got: %s", toJSON(tc.ret), toJSON(result))
			}
*/
		})
	}
}
