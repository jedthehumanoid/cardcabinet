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

// Card is a text with properties, meant for displaying on a board.
type Card struct {
	Title       string                 `json:"title"`
	Contents    string                 `json:"contents"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	Frontmatter string                 `json:"frontmatter,omitempty"`
}

func ReadDir(dir string) ([]Card, []Board) {

	files := findFiles(dir)

	cards := ReadCards(files)
	boards := ReadBoards(files)

	for i, _ := range cards {
		cards[i].Title = strings.TrimPrefix(cards[i].Title, dir)
	}

	for i, _ := range boards {
		boards[i].Title = strings.TrimPrefix(boards[i].Title, dir)
	}

	boards = append(boards, getLabels(cards)...)
	boards = append(boards, getFolders(cards)...)

	// Add nameless board for all the cards
	deck := Deck{}
	for _, card := range cards {
		deck.Cards = append(deck.Cards, card.Title)
	}

	board := Board{}
	board.Decks = append(board.Decks, deck)
	boards = append(boards, board)

	boards = refreshBoards(cards, boards)
	return cards, boards
}

func refreshBoards(cards []Card, boards []Board) []Board {
	for i, board := range boards {
		for i, deck := range board.Decks {
			if len(deck.Labels) > 0 {
				deck.Cards = filterLabels(cards, deck.Labels)
			}
			deck.Cards = filterPath(deck.Cards, board.Title)
			board.Decks[i] = deck
		}
		boards[i] = board
	}
	return boards
}

func getLabels(cards []Card) []Board {
	boards := []Board{}
	labels := []string{}

	for _, card := range cards {
		for _, label := range asStringSlice(card.Properties["labels"]) {
			if !ContainsString(labels, label) {
				labels = append(labels, label)
			}
		}
	}

	for _, label := range labels {
		deck := Deck{}
		deck.Title = label
		deck.Labels = []string{label}

		board := Board{}
		board.Title = "+" + label
		board.Decks = append(board.Decks, deck)

		boards = append(boards, board)
	}
	return boards
}

func GetCard(cards []Card, title string) (Card, error) {
	for _, card := range cards {
		if card.Title == title {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("Cannot find %s", title)
}

func getFolders(cards []Card) []Board {
	boards := []Board{}
	folders := []string{}

	for _, card := range cards {
		folder := filepath.Dir(card.Title)
		if folder != "." && !ContainsString(folders, folder) {
			folders = append(folders, folder)
		}
	}

	for _, folder := range folders {
		deck := Deck{}
		deck.Title = folder
		deck.Path = folder

		board := Board{}
		board.Title = "/" + folder + "/"
		board.Decks = append(board.Decks, deck)

		boards = append(boards, board)
	}

	return boards
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

func ReadCards(files []string) []Card {
	cards := []Card{}

	for _, file := range files {
		if !IsCard(file) {
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

func MarshalFrontmatter(card Card, fences bool) string {
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

func filterLabels(cards []Card, labels []string) []string {
	ret := []string{}
	for _, card := range cards {
		l := asStringSlice(card.Properties["labels"])
		if ContainsStrings(l, labels) {
			ret = append(ret, card.Title)
		}
	}
	return ret
}

func filterPath(cards []string, p string) []string {
	ret := []string{}

	p = filepath.Dir(p)
	p = strings.TrimPrefix(p, "/")
	p = strings.TrimPrefix(p, ".")

	for _, card := range cards {
		if strings.HasPrefix(card, p) {
			ret = append(ret, card)
		}
	}
	return ret
}
