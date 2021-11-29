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
	// Map[Cards]
	for _, card := range cards {
		if card.Match(deck.Filter) {
			ret = append(ret, card)
		}
	}

	return ret
}

// ReadBoards reads all boards in dir into boards
func ReadBoards(dir string, recursive bool) []Board {
	boards := []Board{}
	for _, file := range findFiles(dir, recursive) {
		if !strings.HasSuffix(file, ".board.toml") {
			continue
		}

		board, err := ReadBoard(file)
		if err != nil {
			panic(err)
		}
		boards = append(boards, board)
	}
	for i := range boards {
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
