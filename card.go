package cardcabinet

import (
	"bytes"
	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Card is a text with properties, meant for displaying on a board.
type Card struct {
	Title       string                 `json:"title"`
	Contents    string                 `json:"contents"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	Frontmatter string                 `json:"frontmatter,omitempty"`
}

// ReadCardFile takes a file path, reading file in to a card.
func ReadCard(path string) (Card, error) {
	var card Card

	card.Title = ToSlug(strings.TrimSuffix(path, ".md")) + ".md"

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

func IsCard(file string) bool {
	return strings.HasSuffix(file, ".md")
}

func ReadCards(dir string) []Card {
	cards := []Card{}

	for _, file := range findFiles(dir) {
		if !IsCard(file) {
			continue
		}
		card, err := ReadCard(file)
		if err != nil {
			panic(err)
		}
		card.Title = strings.TrimPrefix(card.Title, dir)
		cards = append(cards, card)
	}
	return cards
}

func MarshalFrontmatter(card Card) string {
	ret := ""

	switch card.Frontmatter {
	case "yaml":
		b, _ := yaml.Marshal(card.Properties)
		frontmatter := strings.TrimSpace(string(b))
		if frontmatter != "{}" {
			ret = "---\n" + frontmatter + "\n---"
		}
	case "toml":
		buf := new(bytes.Buffer)
		toml.NewEncoder(buf).Encode(card.Properties)
		frontmatter := strings.TrimSpace(buf.String())
		if frontmatter != "" {
			ret = "+++\n" + frontmatter + "\n+++"
		}
	}

	return ret
}

func FilterLabels(cards []Card, labels []string) []Card {
	ret := []Card{}
	for _, card := range cards {
		l := asStringSlice(card.Properties["labels"])
		if ContainsStrings(l, labels) {
			ret = append(ret, card)
		}
	}
	return ret
}

func GetLabels(cards []Card) []string {
	labels := []string{}
	for _, card := range cards {
		for _, label := range asStringSlice(card.Properties["labels"]) {
			if !ContainsString(labels, label) {
				labels = append(labels, label)
			}
		}
	}
	return labels
}
