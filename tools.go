package cardcabinet

import (
	"os"
	"path/filepath"
	"strings"
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
