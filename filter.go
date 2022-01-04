package cardcabinet

import (
	"encoding/json"
	"github.com/jedthehumanoid/cardcabinet/rpn"
	"strings"
)

func (card Card) Match(filter string) bool {
	_ = rpn.Split(filter)
	//s := expandStackVariables(stack, card)
	//return rpn.Query(s)
	return true
}
/*
func expandStackVariables(stack []string, card Card) []string {
	ret := []string{}
	// mfr map strings strings
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
*/

func expandFromContext(s string, context string) string {
	var ctx map[string]interface{}

	err := json.Unmarshal([]byte(context), &ctx)
	if err != nil {
		panic(err)
	}

	slice := []string{}
	for _, token := range rpn.Split(s) {
		value, exists := ctx[token]
		if exists {
			slice = append(slice, toJSON(value))
		} else {
			slice = append(slice, token)
		}
	}


	return strings.Join(slice, " ")
}


func toJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}
