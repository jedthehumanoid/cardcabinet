package rpn

import(
	"testing"
)

func TestEvalEmpty(t *testing.T) {
	rpn := Rpn{}

	err := rpn.Eval([]string{})
	if err != nil {
		t.Errorf("expected nil")
	}
}

func TestEvalSingleValue(t *testing.T) {
	rpn := Rpn{}

	err := rpn.Eval([]string{"true"})
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.stack) != 1 {
		t.Errorf("expected stack to have one value")
	}

	if rpn.stack[0] != "true" {
		t.Errorf("expected value to be \"true\"")
	}
}

func TestEvalTwoValues(t *testing.T) {
	rpn := Rpn{}

	err := rpn.Eval([]string{"true", "false"})
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.stack) != 2 {
		t.Errorf("expected stack to have two values")
	}

	if rpn.stack[0] != "true" || rpn.stack[1] != "false"{
		t.Errorf("expected values to be \"true\" and \"false\"")
	}
}

func TestEvalAnd(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval([]string{"true", "false", "&"})
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}
	
	if rpn.stack[0] != "false" {
		t.Errorf("expected result to be false")
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{"true", "true", "&"})
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}
	
	if rpn.stack[0] != "true" {
		t.Errorf("expected result to be true")
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{"false", "false", "&"})
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}
	
	if rpn.stack[0] != "false" {
		t.Errorf("expected result to be false")
	}
}


func TestEvalOr(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval([]string{"true", "false", "|"})
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}
	
	if rpn.stack[0] != "true" {
		t.Errorf("expected result to be true")
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{"true", "true", "|"})
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}
	
	if rpn.stack[0] != "true" {
		t.Errorf("expected result to be true")
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{"false", "false", "|"})
	if err != nil {
		t.Errorf("expected nil")
	}
	if len(rpn.stack) != 1 {
		t.Errorf("expected stack to have one value, %v", rpn)
	}
	
	if rpn.stack[0] != "false" {
		t.Errorf("expected result to be false")
	}
}


func TestEvalContainsWithSlice(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval([]string{`["foo", "bar"]`, `"foo"`, "..."})
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.stack) == 1 && rpn.stack[0] == "true") {
		t.Errorf("expected stack to have true, had %v", rpn.stack)
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{`["foo", "bar"]`, `"baz"`, "..."})
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.stack) == 1 && rpn.stack[0] == "false") {
		t.Errorf("expected stack to have false, had %v", rpn.stack)
	}
}

func TestEvalContainsWithString(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval([]string{`"foo"`, `"oo"`, "..."})
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.stack) == 1 && rpn.stack[0] == "true") {
		t.Errorf("expected stack to have true, had %v", rpn.stack)
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{`"foo"`, `"oof"`, "..."})
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.stack) == 1 && rpn.stack[0] == "false") {
		t.Errorf("expected stack to have false, had %v", rpn.stack)
	}
}


func TestEvalEqual(t *testing.T) {
	rpn := Rpn{}
	err := rpn.Eval([]string{`"foo"`, `"foo"`, "="})
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.stack) == 1 && rpn.stack[0] == "true") {
		t.Errorf("expected stack to have true, had %v", rpn.stack)
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{`"foo"`, `"oof"`, "="})
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.stack) == 1 && rpn.stack[0] == "false") {
		t.Errorf("expected stack to have false, had %v", rpn.stack)
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{"2", "2", "="})
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.stack) == 1 && rpn.stack[0] == "true") {
		t.Errorf("expected stack to have true, had %v", rpn.stack)
	}

	rpn = Rpn{}
	err = rpn.Eval([]string{"2", "3", "="})
	if err != nil {
		t.Errorf("expected nil")
	}
	if !(len(rpn.stack) == 1 && rpn.stack[0] == "false") {
		t.Errorf("expected stack to have false, had %v", rpn.stack)
	}

}

func TestSplit(t *testing.T) {
	result := Split(`split here but "not here" and 'not here'`)
	if toJSON(result) != `["split","here","but","\"not here\"","and","'not here'"]` {
		t.Errorf("Unexpected result, got: \n%s", toJSON(result))
	}
	
}