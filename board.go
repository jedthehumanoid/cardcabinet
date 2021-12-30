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
	Name   string `toml:"name" json:"name"`
	Filter string `toml:"filter" json:"filter"`
}

// Path returns path of board
func (board Board) Path() string {
	return getPath(board.Name)
}

// FilterCards returns all cards that matches deck filter
func (deck Deck) FilterCards(cards []Card) []Card {
	ret := []Card{}
	// mfr filter[Cards]
	for _, card := range cards {
		if card.Match(deck.Filter) {
			ret = append(ret, card)
		}
	}

	return ret
}

func isBoard(filename string) bool {
	return strings.HasSuffix(filename, ".board.toml")
}

func FilterStrings(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// ReadBoards reads all boards in dir into boards
func ReadBoards(dir string, recursive bool) []Board {
	boards := []Board{}
	for _, file := range FilterStrings(findFiles(dir, recursive), isBoard) {
		board, err := ReadBoard(file)
		if err != nil {
			panic(err)
		}
		boards = append(boards, board)
	}
	for i := range boards {
		// mfr Map cards cards
		boards[i].Name = strings.TrimPrefix(boards[i].Name, dir)
		boards[i].Name = strings.TrimPrefix(boards[i].Name, "/")
	}
	return boards
}

// ReadBoard reads a board.toml file into a board
func ReadBoard(filename string) (Board, error) {
	var board Board
	board.Name = strings.TrimSuffix(filename, ".board.toml")
	contents, err := ioutil.ReadFile(filepath.FromSlash(filename))
	if err != nil {
		return board, err
	}
	_, err = toml.Decode(string(contents), &board)
	return board, err
}
