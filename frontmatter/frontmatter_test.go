package frontmatter

import (
	"encoding/json"
	"testing"
)

func TestHasFrontmatter(t *testing.T) {
	result := HasFrontmatter("---\nfrontmatter\n---\ncontents")
	if result != "YAML" {
		t.Errorf("Result shoud be YAML, is: %s", result)
	}

	result = HasFrontmatter("+++\nfrontmatter\n+++\ncontents")
	if result != "TOML" {
		t.Errorf("Result shoud be TOML, is: %s", result)
	}

	result = HasFrontmatter("##\nfrontmatter\n#@\ncontents")
	if result != "" {
		t.Errorf("Result shoud be empty, is: %s", result)
	}
}

func TestRemoveFrontmatter(t *testing.T) {
	result := RemoveFrontmatter("---\nfrontmatter\n---\ncontents")
	if result != "contents" {
		t.Errorf("Unexpected result: %s", result)
	}

	result = RemoveFrontmatter("+++\nfrontmatter\n+++\ncontents")
	if result != "contents" {
		t.Errorf("Unexpected result: %s", result)
	}

	result = RemoveFrontmatter("##\nfrontmatter\n#@\ncontents")
	if result != "##\nfrontmatter\n#@\ncontents" {
		t.Errorf("Unexpected result: %s", result)
	}
}

func TestMarshalFrontmatter(t *testing.T) {

	i := map[string]string{
		"status": "Done",
	}
	result := MarshalFrontmatter(i, "YAML", false)
	expected := "status: Done"
	if result != expected {
		t.Errorf("Unexpected result:\n%s\nexpected:\n%s\n", result, expected)
	}

	result = MarshalFrontmatter(i, "YAML", true)
	expected = "---\nstatus: Done\n---"
	if result != expected {
		t.Errorf("Unexpected result:\n%s\nexpected:\n%s\n", result, expected)
	}

	result = MarshalFrontmatter(i, "TOML", false)
	expected = "status = \"Done\""
	if result != expected {
		t.Errorf("Unexpected result:\n%s\nexpected:\n%s\n", result, expected)
	}

	result = MarshalFrontmatter(i, "TOML", true)
	expected = "+++\nstatus = \"Done\"\n+++"
	if result != expected {
		t.Errorf("Unexpected result:\n%s\nexpected:\n%s\n", result, expected)
	}

}

func TestUnmarshalFrontmatter(t *testing.T) {

	expected := map[string]string{
		"status": "Done",
	}
	result, _ := UnmarshalFrontmatter("---\nstatus: Done\n---")
	if toJSON(result) != toJSON(expected) {
		t.Errorf("Unexpected result:\n%s\nexpected:\n%s\n", result, expected)
	}
	result, _ = UnmarshalFrontmatter("+++\nstatus = \"Done\"\n+++")
	if toJSON(result) != toJSON(expected) {
		t.Errorf("Unexpected result:\n%s\nexpected:\n%s\n", result, expected)
	}

	result, _ = UnmarshalFrontmatter("##\nasfjsf\n@@\n")
	if toJSON(result) != toJSON(nil) {
		t.Errorf("Unexpected result:\n%s\nexpected:\n%s\n", result, expected)
	}

}

func toJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}
