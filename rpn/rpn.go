package rpn

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type Rpn struct {
	Variables map[string]interface{}
	Stack     []string
}

func (rpn *Rpn) Eval(query string) error {
	fmt.Println("EVAL:", query)
	fmt.Println("Variables: ", rpn.Variables)
	for i, s := range Split(query) {
		fmt.Printf("%d: %v -- %s\n", i, rpn.Stack, s)
		switch s {
		case "...":
			rpn.contains()
		case "=":
			rpn.equals()
		case "&":
			rpn.and()
		case "|":
			rpn.or()
		default:
			rpn.push(s)
		}
	}
	fmt.Println("FINISHED:", rpn.Stack)
	return nil
}

func ExpandFromContext(s string, context string) string {
	var ctx map[string]interface{}

	err := json.Unmarshal([]byte(context), &ctx)
	if err != nil {
		panic(err)
	}

	slice := []string{}
	for _, token := range Split(s) {
		value, exists := ctx[token]
		if exists {
			slice = append(slice, toJSON(value))
		} else {
			slice = append(slice, token)
		}
	}

	return strings.Join(slice, " ")
}

// typeOf only exists because of type switching a slice returns []interface{}
func typeOf(i interface{}) string {
	switch t := i.(type) {
	case []interface{}:
		return fmt.Sprintf("[]%T", t[0])
	default:
		return fmt.Sprintf("%T", i)
	}
}

func (rpn *Rpn) pop() (string, error) {
	len := len(rpn.Stack)
	val := rpn.Stack[len-1]
	rpn.Stack = rpn.Stack[:len-1]
	return val, nil 
}

func (rpn *Rpn) popVal() (interface{}, error) {
	j, err := rpn.pop()
	if err != nil {
		return nil, err
	}
	return fromJSON(j), nil

}

func (rpn *Rpn) push(value string) error {
	rpn.Stack = append(rpn.Stack, value)
	return nil
}

func (rpn *Rpn) expand(s string) (string, error) {
	fmt.Println("expand:", s)
	value, exists := rpn.Variables[s]
	if exists {
		fmt.Println(value)
		return toJSON(value), nil
	}
	fmt.Println("missing")
	return "", fmt.Errorf("Missing variable: %s", s)
}

func (rpn *Rpn) contains() error {
	val, _ := rpn.pop()
	if fromJSON(val) == nil { 
		val, _ = rpn.expand(val)
	}
	b := fromJSON(val)

	val, _ = rpn.pop()
	if fromJSON(val) == nil { 
		val, _ = rpn.expand(val)
	}
	a := fromJSON(val)

	switch typeOf(a) {
	case "string":
		rpn.push(fmt.Sprintf("%t", strings.Contains(a.(string), b.(string))))
	case "[]string":
		rpn.push(fmt.Sprintf("%t", containsString(asStringSlice(a), b.(string))))
	default:
		rpn.push("false")
	}
	return nil
}

func (rpn *Rpn) equals() error {
	b, _ := rpn.pop()
	if fromJSON(b) == nil {
		b, _ = rpn.expand(b)
	}
	a, _ := rpn.pop()
	if fromJSON(a) == nil {
		a, _ = rpn.expand(a)
	}
	
	rpn.push(fmt.Sprintf("%t", a == b))
	return nil
}

func (rpn *Rpn) or() error {
	b, _ := rpn.popVal()
	a, _ := rpn.popVal()
	rpn.push(fmt.Sprintf("%t", a.(bool) || b.(bool)))
	return nil
}

func (rpn *Rpn) and() error {
	b, _ := rpn.popVal()
	a, _ := rpn.popVal()
	rpn.push(fmt.Sprintf("%t", a.(bool) && b.(bool)))
	return nil
}

func Split(s string) []string {
	r := regexp.MustCompile(`[^\s"'[]+|"[^"]*"|'[^']*'|\[[^\]]*\]`)
	ret := r.FindAllString(s, -1)
	return ret
}

func fromJSON(j string) interface{} {
	var i interface{}
	json.Unmarshal([]byte(j), &i)
	return i
}

func toJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}
func asStringSlice(i interface{}) []string {
	if i == nil {
		return []string{}
	}
	ret := []string{}
	// mfr map interface strings
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
