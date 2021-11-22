package cardcabinet

import (
	"strings"
	"regexp"
	"fmt"
	"encoding/json"
)

func _typeOf(i interface{}) string {
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


func QueryRPN(card Card, query []string) bool {
	operators :=  map[string]func(Card, []string) []string {
		"...": contains,
		"=": equals,
		"&&": and,
		"||": or,
	}
	
	stack := []string{}
	if len(query) == 0 {
		return true
	}
	for _,s := range query {
	
		function, operator := operators[s]
		if !operator {
			push(&stack, s)
		} else {
			stack = function(card, stack)
		}		
	}
	return len(stack) == 1 && stack[0] == "true"
}

func contains(card Card, stack []string) []string {
	b := pop(&stack)
	a := pop(&stack)
	return append(stack, fmt.Sprintf("%t", ContainsString(asStringSlice(card.Properties[a]), b)))
}
	
func equals(card Card, stack []string) []string {
	b := pop(&stack)
	a := pop(&stack)
	if a == "name" {
		return append(stack, fmt.Sprintf("%t", card.Name == b))
	} else {
		return append(stack, fmt.Sprintf("%t", card.Properties[a] == b))
	}	
}
	
func or(card Card, stack []string) []string {
	a := pop(&stack)
	b := pop(&stack)
	if a == "true" || b == "true"  {
		return append(stack, "true")	
	} else {
		return append(stack, "false")		
	}
}


func and(card Card, stack []string) []string {
	a := pop(&stack)
	b := pop(&stack)
	
	if a == "true" && b == "true"  {
		return append(stack, "true")	
	} else {
		return append(stack, "false")	
	}
}

// Solution from https://stackoverflow.com/questions/47489745/splitting-a-string-at-space-except-inside-quotation-marks
func split(s string) []string {
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`) 
	ret := r.FindAllString(s, -1)
	for i, _ := range ret {
	ret[i] = strings.Trim(ret[i], "\"'")
	}
	return ret
}

func QueryCards(cards []Card, querystring string) []Card {
	ret := []Card{}
	qs := split(querystring)
	for _, card := range cards {
		if QueryRPN(card, qs) {
			ret = append(ret, card)
		}
	}
	return ret
}

func toJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func pop(slice *[]string) string {
   length := len(*slice)
   ret := (*slice)[length-1]
   *slice = append((*slice)[:length-1])
   return ret
}

func push(slice *[]string, value string) {
   *slice = append(*slice, value)
}

func prettystack(stack []string) string {
	ret := []string{}
	for _, s := range stack {
		if strings.Contains(s, " ") {
			s = "'" + s + "'"
		}
		ret = append(ret, s)
	}
	return strings.Join(ret, " ")

}
