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
	Name  string `toml:"name" json:"name"`
	Query string `toml:"query" json:"query"`
}

// Path returns path of board
func (board Board) Path() string {
	dir := filepath.Dir(board.Name)
	dir = strings.TrimPrefix(dir, ".") + "/"
	if dir == "//" || dir == "/" {
		dir = ""
	}
	return dir
}

// Get returns all cards in deck
func (deck Deck) Get(cards []Card) []Card {
	cards = QueryCards(cards, deck.Query)
	return cards
}

// ReadBoards reads all boards in dir into boards
func ReadBoards(dir string, recursive bool) []Board {
	boards := []Board{}
	for _, file := range FindFiles(dir, recursive) {
		if !strings.HasSuffix(file, ".board.toml") {
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

// ReadBoard reads a board.toml file into a board
func ReadBoard(dir string) (Board, error) {
	var board Board
	board.Name = strings.TrimSuffix(dir, ".board.toml")
	contents, err := ioutil.ReadFile(filepath.FromSlash(dir))
	if err != nil {
		return board, err
	}
	_, err = toml.Decode(string(contents), &board)
	return board, err
}
