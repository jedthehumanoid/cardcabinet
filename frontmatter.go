package cardcabinet

import (
	"regexp"
)

var yamlregexp = regexp.MustCompile("(?ms)^---($.*)^---$")
var tomlregexp = regexp.MustCompile("(?ms)^\\+\\+\\+($.*)^\\+\\+\\+$")

// GetFrontmatter returns type of frontmatter, full match (including fences), and inside of match
func GetFrontmatter(b []byte) (string, []byte, []byte) {
	match := yamlregexp.FindSubmatch(b)
	if len(match) > 1 {
		return "yaml", match[0], match[1]
	}
	match = tomlregexp.FindSubmatch(b)
	if len(match) > 1 {
		return "toml", match[0], match[1]
	}

	return "", []byte{}, []byte{}
}
