package fsutil

import (
	"io/fs"

	"github.com/ejuju/go-sveltekit/pkg/logutil"
)

// Info retrieves information about a file system
func LogFiles(l logutil.Logger, lvl logutil.LogLevel, fsys fs.FS) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		output := ""

		if d.IsDir() {
			return nil
		}

		output += "Found file: " + path
		return l.Log(logutil.LogLevelDebug, output)
	})
}
