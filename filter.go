package cardcabinet

import "fmt"

type filter struct {
	Attribute string
	Method string
	RHS string
}

func Filter(cards []Card, s string) []Card {
	fmt.Println("filter!")
	return []Card{}
}

func parseFilter(s string) [][]filter{

}
// Ett filter skulle kunna se ut:
// labels contains "jobb"
// labels contains ["jobb","task"]
// Kan man json-parsa just tredje delen
// labels has jobb and ( assignee = Mattias Jadelius or assignee = Anders Englund )
// labels includes jobb,task and assignee = Mattias Jadelius or labels has jobb and assignee = Anders Englund
