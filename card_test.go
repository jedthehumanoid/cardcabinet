package cardcabinet

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
	//	"fmt"
)

func TestReadCard(t *testing.T) {

	tt := []struct {
		Input  string
		Output string
	}{
		{"testdata/card.md", "testdata/card.json"},
		{"testdata/wonky-middlematter.md", "testdata/wonky-middlematter.json"},
	}

	for _, tc := range tt {
		t.Run(tc.Input, func(t *testing.T) {
			card, _ := ReadCard(tc.Input)
			result := toJSONIndent(card)
			b, _ := ioutil.ReadFile(tc.Output)
			expected := strings.TrimSpace(string(b))

			if result != expected {
				t.Errorf("unexpected result, got:\n%s\n, expected:\n%s\n", result, expected)
			}
		})
	}
}

func TestReadCards(t *testing.T) {
	cards := ReadCards("testdata/", false)

	c := []string{}
	for _, card := range cards {
		c = append(c, card.Name)
	}
	result := toJSON(c)

	expected := toJSON([]string{"add-magnets.md", "card.md", "types.md", "wonky-middlematter.md"})

	if result != expected {
		t.Errorf("unexpected result, got: %s, expected: %s", result, expected)
	}
}


func toJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func toJSONIndent(i interface{}) string {
	b, _ := json.MarshalIndent(i, "", "   ")
	return string(b)
}


func TestMatch(t *testing.T) {
	tt := []struct {
		query string
		card  string
		match bool
	}{
		{`labels "todo" ... labels "jobb" ... |`, "add-magnets.md", true},
		{`labels "todo"... labels "test" ... &`, "add-magnets.md", true},
		{`assignee "Ture Sventton" =`, "add-magnets.md", true},
	}

	for _, tc := range tt {
		t.Run(tc.query, func(t *testing.T) {
			card, err := ReadCard("testdata/" + tc.card)
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
