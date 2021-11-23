package rpn

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var Operators = map[string]func([]string) []string{
	"...": contains,
	"=":   equals,
	"&&":  and,
	"||":  or,
}

func Query(query []string) bool {
	stack := []string{}
	if len(query) == 0 {
		return true
	}
	for _, s := range query {
		function, operator := Operators[s]
		if !operator {
			push(&stack, s)
		} else {
			stack = function(stack)
		}
	}
	return len(stack) == 1 && stack[0] == "true"
}

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

func pop(slice *[]string) string {
	length := len(*slice)
	ret := (*slice)[length-1]
	*slice = append((*slice)[:length-1])
	return ret
}

func push(slice *[]string, value string) {
	*slice = append(*slice, value)
}

func contains(stack []string) []string {
	b := fromJSON(pop(&stack))
	a := fromJSON(pop(&stack))
	return append(stack, fmt.Sprintf("%t", containsString(asStringSlice(a), b.(string))))
}

func equals(stack []string) []string {
	b := pop(&stack)
	a := pop(&stack)
	return append(stack, fmt.Sprintf("%t", a == b))

}

func or(stack []string) []string {
	b := pop(&stack)
	a := pop(&stack)
	return append(stack, fmt.Sprintf("%t", a == "true" || b == "true"))
}

func and(stack []string) []string {
	b := pop(&stack)
	a := pop(&stack)
	return append(stack, fmt.Sprintf("%t", a == "true" && b == "true"))
}

// Solution from https://stackoverflow.com/questions/47489745/splitting-a-string-at-space-except-inside-quotation-marks
func Split(s string) []string {
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)
	ret := r.FindAllString(s, -1)
	return ret
}

func fromJSON(j string) interface{} {
	var i interface{}
	json.Unmarshal([]byte(j), &i)
	return i
}

func asStringSlice(i interface{}) []string {
	ret := []string{}
	if i == nil {
		return ret
	}
	for _, v := range i.([]interface{}) {
		ret = append(ret, v.(string))
	}
	return ret
}

// ContainsString searches slice for string
func containsString(list []string, s string) bool {
	for _, val := range list {
		if val == s {
			return true
		}
	}
	return false
}
