package frontmatter

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v2"
)

var yamlregexp = regexp.MustCompile(`(?ms)\A\s*?---([\s\S]*?)^---$`)
var tomlregexp = regexp.MustCompile(`(?ms)\A\s*?\+\+\+([\s\S]*?)^\+\+\+$`)

func HasFrontmatter(s string) string {
	match := yamlregexp.FindStringSubmatch(s)
	if len(match) > 1 {
		return "YAML"
	}
	match = tomlregexp.FindStringSubmatch(s)
	if len(match) > 1 {
		return "TOML"
	}
	return ""

}

func UnmarshalFrontmatter(s string) (map[string]interface{}, error) {

	var ret map[string]interface{}

	switch HasFrontmatter(s) {
	case "YAML":
		match := yamlregexp.FindStringSubmatch(s)
		err := yaml.Unmarshal([]byte(match[1]), &ret)
		return ret, err
	case "TOML":
		match := tomlregexp.FindStringSubmatch(s)
		_, err := toml.Decode(match[1], &ret)
		return ret, err
	}
	return ret, nil
}

func RemoveFrontmatter(s string) string {
	switch HasFrontmatter(s) {
	case "YAML":
		match := yamlregexp.FindStringSubmatch(s)
		s = strings.TrimPrefix(s, match[0])
	case "TOML":
		match := tomlregexp.FindStringSubmatch(s)
		s = strings.TrimPrefix(s, match[0])
	}
	return strings.TrimSpace(s)
}

func MarshalFrontmatter(i interface{}, frontmattertype string, fences bool) string {
	ret := ""

	switch frontmattertype {
	case "YAML":
		b, _ := yaml.Marshal(i)
		frontmatter := strings.TrimSpace(string(b))
		if frontmatter != "{}" {
			ret = frontmatter
			if fences {
				ret = "---\n" + ret + "\n---"
			}
		}
	case "TOML":
		buf := new(bytes.Buffer)
		toml.NewEncoder(buf).Encode(i)
		frontmatter := strings.TrimSpace(buf.String())
		if frontmatter != "" {
			ret = frontmatter
			if fences {
				ret = "+++\n" + ret + "\n+++"
			}
		}
	}
	return ret
}
