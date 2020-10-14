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

func Filter(cards []Card, s string) []Card {
	fmt.Println("filter!")
	return []Card{}
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
