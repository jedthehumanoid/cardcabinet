package rpn

import (
	"testing"
)

func TestEvalEmpty(t *testing.T) {
	rpn := Rpn{}

	err := rpn.Eval("")
	if err != nil {
		t.Errorf("expected nil")
	}
}

func TestEvalSingleValue(t *testing.T) {
	rpn := Rpn{}

	err := rpn.Eval("true")
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.Stack) != 1 {
		t.Errorf("expected stack to have one value")
	}

	if rpn.Stack[0] != "true" {
		t.Errorf("expected value to be \"true\"")
	}
}

func TestEvalTwoValues(t *testing.T) {
	rpn := Rpn{}

	err := rpn.Eval("true false")
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.Stack) != 2 {
		t.Errorf("expected stack to have two values")
	}

	if rpn.Stack[0] != "true" || rpn.Stack[1] != "false" {
		t.Errorf("expected values to be \"true\" and \"false\"")
	}
}

func TestEvalAnd(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval("true false &")
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.Stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}

	if rpn.Stack[0] != "false" {
		t.Errorf("expected result to be false")
	}

	rpn = Rpn{}
	err = rpn.Eval("true true &")
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.Stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}

	if rpn.Stack[0] != "true" {
		t.Errorf("expected result to be true")
		t.Errorf("%v", rpn.Stack)
	}

	rpn = Rpn{}
	err = rpn.Eval("false false &")
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.Stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}

	if rpn.Stack[0] != "false" {
		t.Errorf("expected result to be false")
	}
}

func TestEvalOr(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval("true false |")
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.Stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}

	if rpn.Stack[0] != "true" {
		t.Errorf("expected result to be true")
	}

	rpn = Rpn{}
	err = rpn.Eval("true true |")
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.Stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}

	if rpn.Stack[0] != "true" {
		t.Errorf("expected result to be true")
	}

	rpn = Rpn{}
	err = rpn.Eval("false false |")
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.Stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}

	if rpn.Stack[0] != "false" {
		t.Errorf("expected result to be false")
	}
}

func TestEvalContainsWithSlice(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval(`["foo", "bar"] "foo" ...`)
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.Stack) == 1 && rpn.Stack[0] == "true") {
		t.Errorf("expected stack to have true, had %v", rpn.Stack)
	}

	rpn = Rpn{}
	err = rpn.Eval(`["foo", "bar"] "baz" ...`)
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.Stack) == 1 && rpn.Stack[0] == "false") {
		t.Errorf("expected stack to have false, had %v", rpn.Stack)
	}
}

func TestEvalContainsWithString(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval(`"foo" "oo" ...`)
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.Stack) == 1 && rpn.Stack[0] == "true") {
		t.Errorf("expected stack to have true, had %v", rpn.Stack)
	}

	rpn = Rpn{}
	err = rpn.Eval(`"foo" "oof" ...`)
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.Stack) == 1 && rpn.Stack[0] == "false") {
		t.Errorf("expected stack to have false, had %v", rpn.Stack)
	}
}

func TestEvalEqual(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval(`"foo" "foo" =`)
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.Stack) == 1 && rpn.Stack[0] == "true") {
		t.Errorf("expected stack to have true, had %v", rpn.Stack)
	}

	rpn = Rpn{}
	err = rpn.Eval(`"foo" "oof" =`)
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.Stack) == 1 && rpn.Stack[0] == "false") {
		t.Errorf("expected stack to have false, had %v", rpn.Stack)
	}

	rpn = Rpn{}
	err = rpn.Eval("2 2 =")
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.Stack) == 1 && rpn.Stack[0] == "true") {
		t.Errorf("expected stack to have true, had %v", rpn.Stack)
	}

	rpn = Rpn{}
	err = rpn.Eval("2 3 =")
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.Stack) == 1 && rpn.Stack[0] == "false") {
		t.Errorf("expected stack to have false, had %v", rpn.Stack)
	}

}

func TestSplit(t *testing.T) {
	result := Split(`split here but "not here" and 'not here'`)
	if toJSON(result) != `["split","here","but","\"not here\"","and","'not here'"]` {
		t.Errorf("Unexpected result, got: \n%s", toJSON(result))
	}

}

func TestPush(t *testing.T) {
	rpn := Rpn{}
	rpn.push("foo")
	if len(rpn.Stack) != 1 || rpn.Stack[0] != "foo" {
		t.Errorf("Expected exactly one value: \"foo\"")
	}
	rpn.push("bar")
	if len(rpn.Stack) != 2 || rpn.Stack[0] != "foo" || rpn.Stack[1] != "bar" {
		t.Errorf("Expected two values, \"foo\" and \"bar\"")
	}

}

func TestPop(t *testing.T) {
	rpn := Rpn{}
	rpn.push(`"foo"`)
	rpn.push(`"bar"`)

	val, err := rpn.pop()
	if err != nil {
		panic(err)
	}
	if len(rpn.Stack) != 1 || rpn.Stack[0] != `"foo"` {
		t.Errorf("Expected exactly one value: \"foo\"")
	}
	if val != `"bar"` {
		t.Errorf("%s", val)
		t.Errorf("Expected to have popped \"bar\"")
	}

	val, err = rpn.pop()
	if err != nil {
		panic(err)
	}
	if len(rpn.Stack) != 0 {
		t.Errorf("Expected zero values, no less, no more")
	}
	if val != `"foo"` {
		t.Errorf("Expected to have popped \"foo\"")
	}

}

func TestPopVal(t *testing.T) {
	rpn := Rpn{}
	rpn.push(`"foo"`)
	rpn.push(`"bar"`)

	val, err := rpn.popVal()
	if err != nil {
		panic(err)
	}

	if val != "bar" {
		t.Errorf("%s", val)
		t.Errorf("Expected to have popped \"bar\"")
	}
}

func TestExpandFromContext(t *testing.T) {

	context := `{
		"name": "jack",
		"labels": ["dull", "boy"]
	}`

	result := ExpandFromContext(`name "jack" = labels "dull" boy ... "name" "labels"`, context)
	expected := `"jack" "jack" = ["dull","boy"] "dull" boy ... "name" "labels"`

	if result != expected {
		t.Errorf("unexpected value: %s, expected: %s", result, expected)
	}
}
