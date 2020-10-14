package cardcabinet

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReadCard(t *testing.T) {
	card, _ := ReadCard("testdata/card.md")
	s, _ := json.MarshalIndent(card, "", "   ")
	fmt.Println(string(s))
	//	t.Errorf("Abs(-1) = %d; want 1", got)

}
