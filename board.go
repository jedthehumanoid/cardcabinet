package cardcabinet

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"path/filepath"
	"strings"
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

func (deck Deck) Get(cards []Card) []string {
	ret := []string{}

	if len(deck.Names) > 0 {
		return deck.Names
	}

	if len(deck.Labels) > 0 {
		ret = append(ret, filterLabels(cards, deck.Labels)...)
	}

	return ret

}

func IsBoard(file string) bool {
	return strings.HasSuffix(file, ".board.toml")
}

func ReadBoards(files []string) []Board {
	boards := []Board{}

	for _, file := range files {
		if !IsBoard(file) {
			continue
		}

		board, err := ReadBoard(file)
		if err != nil {
			panic(err)
		}
		boards = append(boards, board)
	}

	return boards
}

func ReadBoard(path string) (Board, error) {
	var board Board

	board.Name = ToSlug(strings.TrimSuffix(path, "board.toml"))

	contents, err := ioutil.ReadFile(filepath.FromSlash(path))
	if err != nil {
		return board, err
	}

	_, err = toml.Decode(string(contents), &board)

	return board, err
}

func GetBoard(boards []Board, board string) Board {
	for _, b := range boards {
		if b.Name == board {
			return b
		}
	}
	return Board{}
}
