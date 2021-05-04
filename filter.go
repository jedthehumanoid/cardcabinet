package cardcabinet

import (
	"strings"
	"regexp"
	"fmt"
	"encoding/json"
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


func queryRPN(card Card, query []string) bool {
	operators :=  map[string]func([]string) []string {
		"...": contains,
		"=": contains,
		"&&": and,
		"||": or,
	}
	fmt.Println(toJSON(query))

	stack := []string{}
	
	for _,s := range query {
	
		function, operator := operators[s]
		if !operator {
			push(&stack, s)
		} else {
			fmt.Printf("%s? %s\n", s, toJSON(stack))
			stack = function(stack)
		}		
	}
	fmt.Println(stack)
	fmt.Println()
	return false
}

func contains(stack []string) []string {
	_ = pop(&stack)
	_ = pop(&stack)
	return append(stack, "true")
}

func or(stack []string) []string {
	a := pop(&stack)
	b := pop(&stack)
	
	if a == "true" || b == "true"  {
		return append(stack, "true")	
	} else {
		return append(stack, "false")		
	}
}


func and(stack []string) []string {
	a := pop(&stack)
	b := pop(&stack)
	
	if a == "true" && b == "true"  {
		return append(stack, "true")	
	} else {
		return append(stack, "false")	
	}
}

func query(card Card, querystring string) bool {
	querystring = strings.TrimSpace(querystring)
	tokens := strings.Split(querystring, " ")
	property := tokens[0]
	method := tokens[1]
	rhs := strings.Join(tokens[2:], " ")

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
		if typeOf(card.Properties[property]) == "stringSlice" {
			return ContainsString(asStringSlice(card.Properties[property]), rhs)
		}
	case "EQUALS", "=":
		if typeOf(card.Properties[property]) == "string" {
			return card.Properties[property] == rhs
		}
	}
	return false
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

func queryCard(card Card, querystring string) bool {
	keep := false
	for _, x := range strings.Split(querystring, "OR") {
		for _, y := range strings.Split(x, "AND") {
			keep = query(card, y)
			if !keep {
				break
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
