package fsutil

import (
	"fmt"
	"io/fs"
)

// Info retrieves information about a file system
func PrintFiles(fsys fs.FS) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		output := path
		if d.IsDir() {
			output += "/"
		}
		_, err = fmt.Println(output)
		return err
	})
}
