package cardcabinet

import (
	"bytes"
	"fmt"
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

func GetFolders(cards []Card) []string {
	folders := []string{}

	for _, card := range cards {
		folder := filepath.Dir(card.Name)
		if folder != "." && !ContainsString(folders, folder) {
			folders = append(folders, folder)
		}
	}

	return folders
}

func GetCard(cards []Card, name string) (Card, error) {
	for _, card := range cards {
		if card.Name == name {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("Cannot find %s", name)
}

func isCard(file string) bool {
	return strings.HasSuffix(file, ".md")
}

// ReadCardFile takes a file path, reading file in to a card.
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

func ReadCards(files []string) []Card {
	cards := []Card{}

	for _, file := range files {
		if !isCard(file) {
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

func filterLabels(cards []Card, labels []string) []string {
	ret := []string{}
	for _, card := range cards {
		l := asStringSlice(card.Properties["labels"])
		if ContainsStrings(l, labels) {
			ret = append(ret, card.Name)
		}
	}
	return ret
}
