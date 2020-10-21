package cardcabinet

import "testing"

func TestParseFilter(t *testing.T) {
	parseFilter("labels = tasks AND labels = test OR labels = jobb")

}

func TestContains(t *testing.T) {
	card, _ := ReadCard("testdata/types.md")
	Contains(card, "labels", "test")
	Contains(card, "labelsstring", "test")
	Contains(card, "labelsintslice", "test")
}
