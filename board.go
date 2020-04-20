package cardcabinet

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Board struct {
	Title string `toml:"title" json:"title"`
	Decks []Deck `toml:"decks" json:"decks"`
}

type Deck struct {
	Title  string   `toml:"title" json:"title"`
	Labels []string `toml:"labels" json:"labels"`
	Cards  []string `toml:"cards" json:"cards"`
}

func GetCards(cards []Card, board Board) []Deck {
	ret := []string{}

	ret = deck.Cards
	if len(deck.Labels) > 0 {
		ret = filterLabels(cards, deck.Labels)
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

	board.Title = ToSlug(strings.TrimSuffix(path, "board.toml"))

	contents, err := ioutil.ReadFile(filepath.FromSlash(path))
	if err != nil {
		return board, err
	}

	_, err = toml.Decode(string(contents), &board)

	return board, err
}

func GetBoard(boards []Board, board string) Board {
	for _, b := range boards {
		if b.Title == board {
			return b
		}
	}
	return Board{}
}
