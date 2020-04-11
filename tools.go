package cardcabinet

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// FindFiles is like find
func findFiles(path string) []string {
	files := []string{}
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
	return files
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
