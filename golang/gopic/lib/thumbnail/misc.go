package thumbnail

import (
	"os"
	"path/filepath"
)

func createSrc(src string) (*os.File, error) {
	if _, err := os.Stat(src); err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
	}

	f, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// createDstDir creates the path (recursive) to new thumbnail
func createDstDir(dst string) error {
	return os.MkdirAll(filepath.Dir(dst), 0700)
}

func createDst(dst string) (*os.File, error) {
	if err := createDstDir(dst); err != nil {
		return nil, err
	}

	// Create file
	out, err := os.Create(dst)
	if err != nil {
		return out, err
	}

	return out, nil
}
