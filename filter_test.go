package cardcabinet

import "testing"

func TestParseFilter(t *testing.T) {
	cards := ReadCards("testdata/")

	queryCards(cards, "labels ... todo AND labels ... test OR labels ... jobb")

}
