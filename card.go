package cardcabinet

import (
	"github.com/jedthehumanoid/cardcabinet/frontmatter"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

// ReadCard takes a file path, reading file in to a card.
func ReadCard(filename string) (Card, error) {
	var card Card

	card.Name = filename
	contents, err := ioutil.ReadFile(filepath.FromSlash(filename))
	if err != nil {
		return card, err
	}

	card.Contents = string(contents)
	card.Frontmatter = frontmatter.HasFrontmatter(card.Contents)

	properties, err := frontmatter.UnmarshalFrontmatter(card.Contents)
	if err != nil {
		return card, err
	}
	card.Properties = properties
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

// Path returns path of card
func (card Card) Path() string {
	return getPath(card.Name)
}

func getPath(filename string) string {
	dir := filepath.Dir(filename)
	dir = strings.TrimPrefix(dir, ".") + "/"
	if dir == "//" || dir == "/" {
		dir = ""
	}
	return dir
}
