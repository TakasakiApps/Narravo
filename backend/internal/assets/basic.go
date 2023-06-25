package assets

import "github.com/gookit/goutil/fsutil"

func readBytes(path string) []byte {
	if fsutil.FileExists(path) {
		return fsutil.ReadFile(path)
	}
	return nil
}
