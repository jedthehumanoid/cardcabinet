package cardcabinet

import (
	"fmt"
	"reflect"
	"strings"
)

type filter struct {
	Attribute string
	Method    string
	RHS       string
}

func Contains(card Card, property string, s string) bool {
	fmt.Println("contains?")
	fmt.Println(s)
	fmt.Println(card.Properties[property])
	switch t := card.Properties[property].(type) {
	case string:
		fmt.Println("string!")
	case []interface{}:
		fmt.Println("[]interface{}")
		fmt.Println(toJSON(t[0]))
		fmt.Println(reflect.TypeOf(t[0]).String())
		switch t[0].(type) {

		case string:
			fmt.Println("of string")
		case int64:
			fmt.Println("of int")

		}
	}
	fmt.Println()

	return false
}

func parseFilter(s string) [][]filter {
	filters := [][]filter{}
	fmt.Println("FILTER:", s)
	for _, or := range strings.Split(s, "OR") {
		or = strings.TrimSpace(or)
		fmt.Println("OR:", or)
		for _, and := range strings.Split(or, "AND") {
			and = strings.TrimSpace(and)
			fmt.Println("AND:", and)
			tokens := strings.Split(and, " ")
			attr := tokens[0]
			Method := tokens[1]
			rhs := tokens[2:]
			filter := filter{attr, Method, strings.Join(rhs, " ")}
			fmt.Printf("%+v\n", filter)
		}

	}
	return filters
}
