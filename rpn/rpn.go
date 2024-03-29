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

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	ErrStackUnderflow = constError("stack underflow")
	ErrWrongType      = constError("wrong type")
)

func (rpn *Rpn) Eval(query string) error {
	//fmt.Println("EVAL:", query)
	//	fmt.Println("Variables: ", rpn.Variables)
	for _, s := range Split(query) {
		//fmt.Printf("%d: %v -- %s\n", i, rpn.Stack, s)
		switch s {
		case "...":
			rpn.contains()
		case "=":
			rpn.equals()
		case "&", "AND":
			rpn.and()
		case "|", "OR":
			rpn.or()
		case "!", "NOT":
			rpn.not()
		default:
			rpn.push(s)
		}
	}
	//	fmt.Println("FINISHED:", rpn.Stack)
	return nil
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
	if len < 1 {
		return "", ErrStackUnderflow
	}
	val := rpn.Stack[len-1]
	rpn.Stack = rpn.Stack[:len-1]
	return val, nil
}

func (rpn *Rpn) push(value string) error {
	rpn.Stack = append(rpn.Stack, value)
	return nil
}

func (rpn *Rpn) expand(s string) (string, error) {
	value, exists := rpn.Variables[s]
	if exists {
		return toJSON(value), nil
	}
	return "", fmt.Errorf("missing variable: %s", s)
}

func (rpn *Rpn) not() error {
	val, _ := rpn.pop()
	a := fromJSON(val)
	switch a.(type) {
	case bool:
	default:
		return ErrWrongType
	}
	rpn.push(fmt.Sprintf("%t", !a.(bool)))
	return nil

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
	b, err := rpn.pop()
	if err != nil {
		return err
	}
	a, err := rpn.pop()
	if err != nil {
		return err
	}
	if fromJSON(b) == nil {
		b, err = rpn.expand(b)
		if err != nil {
			rpn.push(fmt.Sprintf("%t", false))
			return nil
		}
	}

	if fromJSON(a) == nil {
		a, err = rpn.expand(a)
		if err != nil {
			rpn.push(fmt.Sprintf("%t", false))
			return nil
		}
	}

	rpn.push(fmt.Sprintf("%t", a == b))
	return nil
}

func (rpn *Rpn) or() error {
	val, _ := rpn.pop()
	b := fromJSON(val)
	switch b.(type) {
	case bool:
	default:
		return ErrWrongType
	}
	val, _ = rpn.pop()
	a := fromJSON(val)
	switch a.(type) {
	case bool:
	default:
		return ErrWrongType
	}
	rpn.push(fmt.Sprintf("%t", a.(bool) || b.(bool)))
	return nil
}

func (rpn *Rpn) and() error {
	val, _ := rpn.pop()
	b := fromJSON(val)
	switch b.(type) {
	case bool:
	default:
		return ErrWrongType
	}
	val, _ = rpn.pop()
	a := fromJSON(val)
	switch a.(type) {
	case bool:
	default:
		return ErrWrongType
	}
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
