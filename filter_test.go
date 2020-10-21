package cardcabinet

import "testing"

func TestParseFilter(t *testing.T) {
	cards := ReadCards("testdata/")

	queryCards(cards, "labels contains todo AND labels contains test OR labels contains jobb")

}
