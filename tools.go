package cardcabinet

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	
)

// FindFiles is like find
func FindFiles(path string, recursive bool) []string {
	files := []string{}
	if !recursive {
		dirfiles, _ := ioutil.ReadDir(path)
		for _, file := range dirfiles {
		   if !file.IsDir() {
			files = append(files, file.Name())
		   }
		}
   } else {
	filepath.Walk(path,
		func(file string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !f.IsDir() {
				files = append(files, file)
			}
			return nil
		})
	}
	return files
}

func GetPath(title string) string {
	ret := filepath.Dir(title)
	ret = strings.TrimPrefix(ret, ".")
	if ret != "" {
		ret = ret + "/"
	}
	return ret
}

// FromSlug returns "this format" from "this-format"
func FromSlug(s string) string {
	return strings.Replace(s, "-", " ", -1)
}

// ToSlug returns a "page slug" out of string
// Makes lowercase, removes international accents, turns spaces and
// other nonalphanumeric characters to dashes, but keeps slashes
func ToSlug(s string) string {
	var re = regexp.MustCompile("[^a-z0-9/]+")

	s = strings.ToLower(s)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}), norm.NFC)
	s, _, _ = transform.String(t, s)
	s = re.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}

func asStringSlice(i interface{}) []string {
	ret := []string{}
	if i == nil {
		return ret
	}
	for _, v := range i.([]interface{}) {
		ret = append(ret, v.(string))
	}
	return ret
}

// ContainsString searches slice for string
func ContainsString(list []string, s string) bool {
	for _, val := range list {
		if val == s {
			return true
		}
	}
	return false
}

func ContainsStrings(list []string, ss []string) bool {
	for _, s := range ss {
		if !ContainsString(list, s) {
			return false
		}
	}
	return true
}
