package cardcabinet

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/jedthehumanoid/cardcabinet/frontmatter"
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
func ReadCard(path string) (Card, error) {
	var card Card

	card.Name = path
	contents, err := ioutil.ReadFile(filepath.FromSlash(path))
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

	for _, file := range FindFiles(dir, recursive) {
		if !strings.HasSuffix(file, ".md") {
			continue
		}
		card, err := ReadCard(file)
		if err != nil {
			panic(err)
		}

		cards = append(cards, card)
	}
	return cards
}

// Path returns path of card
func (card Card) Path() string {
	dir := filepath.Dir(card.Name)
	dir = strings.TrimPrefix(dir, ".") + "/"
	if dir == "//" || dir == "/" {
		dir = ""
	}
	return dir
}
