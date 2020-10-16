package cardcabinet

import (
	"bytes"
	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
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
	ret := ""

	switch card.Frontmatter {
	case "yaml":
		b, _ := yaml.Marshal(card.Properties)
		frontmatter := strings.TrimSpace(string(b))
		if frontmatter != "{}" {
			ret = frontmatter
			if fences {
				ret = "---\n" + ret + "\n---"
			}
		}
	case "toml":
		buf := new(bytes.Buffer)
		toml.NewEncoder(buf).Encode(card.Properties)
		frontmatter := strings.TrimSpace(buf.String())
		if frontmatter != "" {
			ret = frontmatter
			if fences {
				ret = "+++\n" + ret + "\n+++"
			}
		}
	}

	return ret
}

// ReadCard takes a file path, reading file in to a card.
func ReadCard(path string) (Card, error) {
	var card Card

	card.Name = path
	contents, err := ioutil.ReadFile(filepath.FromSlash(path))
	if err != nil {
		return card, err
	}
	frontmatter, raw, b := GetFrontmatter(contents)
	card.Frontmatter = frontmatter

	card.Contents = strings.TrimPrefix(string(contents), string(raw))
	card.Contents = strings.TrimSpace(card.Contents)

	switch frontmatter {
	case "yaml":
		err = yaml.Unmarshal(b, &card.Properties)
		return card, err
	case "toml":
		_, err = toml.Decode(string(b), &card.Properties)
		return card, err
	}
	return card, nil
}

// ReadCards reads all cards in directory, and subdirectories
func ReadCards(dir string) []Card {
	cards := []Card{}

	for _, file := range FindFiles(dir) {
		if !strings.HasSuffix(file, ".md") {
			continue
		}
		card, err := ReadCard(file)
		if err != nil {
			panic(err)
		}
		card.Name = strings.TrimPrefix(card.Name, dir)
		cards = append(cards, card)
	}
	return cards
}
