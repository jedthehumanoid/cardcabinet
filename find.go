package cardcabinet

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// FindFiles is like find
func findFiles(dir string, recursive bool) []string {
	files := []string{}
	if !recursive {
		dirfiles, _ := ioutil.ReadDir(dir)
		for _, file := range dirfiles {
			// mfr filter os.fileinfo
			if !file.IsDir() {
				files = append(files, dir+file.Name())
			}
		}
	} else {
		filepath.Walk(dir,
			func(file string, f os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				// mfr filter os.fileinfo
				if !f.IsDir() {
					files = append(files, file)
				}
				return nil
			})
	}
	return files
}
