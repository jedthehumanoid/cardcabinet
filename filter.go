package cardcabinet

import "fmt"

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
	filter := [][]filter{}
	return filter
}
