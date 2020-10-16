package cardcabinet

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Board struct {
	Name  string `toml:"name" json:"name"`
	Decks []Deck `toml:"deck" json:"decks"`
}

type Deck struct {
	Name   string   `toml:"name" json:"name"`
	Labels []string `toml:"labels" json:"labels"`
	Names  []string `toml:"names" json:"names"`
}

func (board Board) Path() string {
	dir := filepath.Dir(board.Name)
	dir = strings.TrimPrefix(dir, ".") + "/"
	if dir == "//" || dir == "/" {
		dir = ""
	}
	return dir
}

func (board Board) Cards(cards []Card) []Card {
	ret := []Card{}
	dir := board.Path()
	for _, card := range cards {
		if strings.HasPrefix(card.Name, dir) {
			card.Name = strings.TrimPrefix(card.Name, dir)
			ret = append(ret, card)
		}
	}
	return ret
}

func (deck Deck) Get(cards []Card) []Card {
	if len(deck.Names) > 0 {
		temp := []Card{}
		for _, card := range cards {
			if ContainsString(deck.Names, card.Name) {
				temp = append(temp, card)
			}
		}
		cards = temp
	} else if len(deck.Labels) > 0 && deck.Labels[0] != "" {
		cards = filterLabels(cards, deck.Labels)
	}

	return cards
}

func ReadBoards(dir string) []Board {
	boards := []Board{}
	for _, file := range FindFiles(dir) {
		if !strings.HasSuffix(file, "board.toml") {
			continue
		}

		board, err := ReadBoard(file)
		if err != nil {
			panic(err)
		}

		board.Name = strings.TrimPrefix(board.Name, dir)
		if board.Name == "" {
			board.Name = "/"
		}
		boards = append(boards, board)
	}
	return boards
}

func ReadBoard(path string) (Board, error) {
	var board Board

	board.Name = strings.TrimSuffix(path, "board.toml")

	contents, err := ioutil.ReadFile(filepath.FromSlash(path))
	if err != nil {
		return board, err
	}

	_, err = toml.Decode(string(contents), &board)

	return board, err
}

func filterLabels(cards []Card, labels []string) []Card {
	ret := []Card{}
	for _, card := range cards {
		l := asStringSlice(card.Properties["labels"])
		if ContainsStrings(l, labels) {
			ret = append(ret, card)
		}
	}
	return ret
}
