package rpn

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type Rpn struct{
	stack []string
}

func (rpn *Rpn) Eval(query []string) error {
	for _, s := range query {
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
	length := len(rpn.stack)
	ret := rpn.stack[length-1]
	 rpn.stack = append(rpn.stack[:length-1])
	return ret, nil
}

func (rpn *Rpn) push(value string) error {
	rpn.stack = append(rpn.stack, value)
	return nil
}

func (rpn *Rpn) contains() error {
	
	val, _ := rpn.pop()
	b := fromJSON(val)
	val, _ = rpn.pop()
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
	b,_ := rpn.pop()
	a,_ := rpn.pop()
	rpn.push(fmt.Sprintf("%t", a == b))
	return nil

}

func (rpn *Rpn) or() error {
	b,_ := rpn.pop()
	a,_ := rpn.pop()
	rpn.push(fmt.Sprintf("%t", a == "true" || b == "true"))
	return nil
}

func (rpn *Rpn) and() error {
	b, _ := rpn.pop()
	a, _ := rpn.pop()
	rpn.push(fmt.Sprintf("%t", a == "true" && b == "true"))
	return nil
}

func Split(s string) []string {
	fmt.Println(s)
	r := regexp.MustCompile(`[^\s"']+|"[^"]*"|'[^']*'`)
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
