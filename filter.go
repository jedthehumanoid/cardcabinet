package cardcabinet

import (
	"encoding/json"
	"github.com/jedthehumanoid/cardcabinet/rpn"
	"strings"
)

func QueryCards(cards []Card, querystring string) []Card {
	ret := []Card{}
	qs := rpn.Split(querystring)
	for _, card := range cards {
		s := expandStackVariables(qs, card)
		if rpn.Query(s) {
			ret = append(ret, card)
		}
	}
	return ret
}

func expandStackVariables(stack []string, card Card) []string {
	ret := []string{}
	for _, token := range stack {
		_, operator := rpn.Operators[token]
		if strings.HasPrefix(token, "\"") || operator {
			ret = append(ret, token)
		} else {
			if token == "name" {
				ret = append(ret, toJSON(card.Name))
			} else {
				ret = append(ret, toJSON(card.Properties[token]))
			}
		}
	}
	return ret
}

func toJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}
