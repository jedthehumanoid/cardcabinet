package cardcabinet

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jedthehumanoid/cardcabinet/frontmatter"
	"github.com/jedthehumanoid/cardcabinet/rpn"
)

// Card is markdown with properties (from frontmatter).
type Card struct {
	Name        string                 `json:"name"`
	Contents    string                 `json:"contents"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	Frontmatter string                 `json:"frontmatter,omitempty"`
}

// MarshalFrontmatter returns properties as frontmatter string
// of correct type, including or excluding fences
func (card Card) MarshalFrontmatter(fences bool) string {
	return frontmatter.MarshalFrontmatter(card.Properties, card.Frontmatter, fences)
}

// Path returns path of card
func (card Card) Path() string {
	return getPath(card.Name)
}

func (card Card) Match(filter string) bool {
	if filter == "" {
		return true
	}
	rpn := rpn.Rpn{}

	// kändes som att det här inverkade på prestandan?
	//context := card.Properties

	//	_, hasname := context["name"]
	//	if !hasname {
	//		context["name"] = card.Name
	//	}
	rpn.Variables = card.Properties
	rpn.Eval(filter)

	if len(rpn.Stack) == 0 {
		return false
	}

	if len(rpn.Stack) == 1 {
		return rpn.Stack[0] == "true"
	}

	// If all values in stack are true, treat it as
	// several queries and default to and between them
	for _, val := range rpn.Stack {
		if val != "true" {
			return false
		}
	}
	return true
}

// ReadCard takes a file path, reading file in to a card.
func ReadCard(filename string) (Card, error) {
	var card Card

	card.Name = filename
	contents, err := ioutil.ReadFile(filepath.FromSlash(filename))
	if err != nil {
		return card, err
	}

	card.Contents = string(contents)
	if strings.Contains(card.Contents, "\r\n") {
		card.Contents = strings.Replace(card.Contents, "\r\n", "\n", -1)
	}
	card.Frontmatter = frontmatter.HasFrontmatter(card.Contents)

	properties, err := frontmatter.UnmarshalFrontmatter(card.Contents)
	if err != nil {
		return card, err
	}
	card.Properties = properties
	if card.Properties == nil {
		card.Properties = map[string]interface{}{}
	}
	card.Contents = frontmatter.RemoveFrontmatter(card.Contents)
	return card, nil
}

// ReadCards reads all cards in directory, and subdirectories
func ReadCards(dir string, recursive bool) []Card {
	cards := []Card{}
	fi, err := os.Stat(dir)
	if err == nil && !fi.Mode().IsDir() {
		return cards
	}
	// mfr Map strings, cards
	for _, file := range findFiles(dir, recursive) {
		// mfr Filter string
		if !strings.HasSuffix(file, ".md") {
			continue
		}
		card, err := ReadCard(file)
		if err != nil {
			panic(err)
		}

		cards = append(cards, card)
	}
	// mfr Map cards, cards
	for i := range cards {
		cards[i].Name = strings.TrimPrefix(cards[i].Name, dir)
		cards[i].Name = strings.TrimPrefix(cards[i].Name, "/")
	}

	return cards
}

func getPath(filename string) string {
	dir := filepath.Dir(filename)
	dir = strings.TrimPrefix(dir, ".") + "/"
	if dir == "//" || dir == "/" {
		dir = ""
	}
	return dir
}
