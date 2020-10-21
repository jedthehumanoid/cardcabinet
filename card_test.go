package cardcabinet

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
	//	"fmt"
)

func TestReadCard(t *testing.T) {
	card, _ := ReadCard("testdata/card.md")

	result := toJSONIndent(card)
	
	b, _ := ioutil.ReadFile("testdata/card.json")
	expected := strings.TrimSpace(string(b))

	if result != expected {
		t.Errorf("unexpected result, got:\n%s\n, expected:\n%s\n", result, expected)
	}
}

func TestReadCards(t *testing.T) {
	cards := ReadCards("testdata/")

	c := []string{}
	for _, card := range cards {
		c = append(c, card.Name)
	}
	result := toJSON(c)
	
	expected := toJSON([]string{"add-magnets.md", "card.md"})

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

