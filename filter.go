package cardcabinet

import (
	"fmt"
	"strings"
)

type filter struct {
	Attribute string
	Method    string
	RHS       string
}

func typeOf(i interface{}) string {
	switch t := i.(type) {
	case string:
		return "string"
	case int64:
		return "int"
	case float64:
		return "float"
	case bool:
		return "bool"
	case []interface{}:
		switch t[0].(type) {
		case string:
			return "stringSlice"
		case int64:
			return "intSlice"
		case float64:
			return "floatSlice"
		case bool:
			return "boolSlice"
		}
	}
	return ""
}

func stringSliceContains() {

}

func query(card Card, querystring string) bool {
	tokens := strings.Split(querystring, " ")
	property := tokens[0]
	method := tokens[1]
	rhs := tokens[2]

	if property == "name" {
		switch method {
		case "CONTAINS", "...":
			return strings.Contains(card.Name, rhs)
		case "EQUALS", "=":
			return card.Name == rhs
		}
	}

	switch method {
	case "CONTAINS", "...":
		if typeOf(property) == "stringSlice" {
			return ContainsString(asStringSlice(property), rhs)
		}
	}
	return false
}

func queryCard(card Card, querystring string) bool {
	for _, or := range strings.Split(querystring, "OR") {
		or = strings.TrimSpace(or)
		keep := true
		for _, and := range strings.Split(or, "AND") {
			and = strings.TrimSpace(and)
			if !query(card, and) {
				keep = false
			}
		}
		if keep {
			return true
		}
	}
	return false
}

func queryCards(cards []Card, querystring string) []Card {
	ret := []Card{}
	for _, card := range cards {
		if queryCard(card, querystring) {
			ret = append(ret, card)
		}
	}
	for _, card := range ret {
		fmt.Println(card.Name)
	}
	return ret
}
