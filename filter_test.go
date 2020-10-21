package cardcabinet

import (
	"testing"
)

func TestQuery(t *testing.T) {

	tt := []struct {
		query string
		cards []string
	}{
		{"labels ... todo AND labels ... test OR labels ... jobb", []string{"add-magnets.md", "types.md"}},
		{"name = types.md", []string{"types.md"}},
		{"assignee = Ture Sventton", []string{"add-magnets.md"}},
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
