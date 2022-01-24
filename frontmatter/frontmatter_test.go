package frontmatter

import (
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
