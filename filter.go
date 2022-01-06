package cardcabinet

import (
	"encoding/json"
	"github.com/jedthehumanoid/cardcabinet/rpn"
)

func (card Card) Match(filter string) bool {
	rpn := rpn.Rpn{}
	rpn.Variables = card.Properties
	rpn.Eval(filter)
	return len(rpn.Stack) == 1 && rpn.Stack[0] == "true"
}

func toJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}
