package files

import (
	"io"
	"os"
)

func Override(path string, content []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	n, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	_, err = f.WriteAt(content, n)
	return err
}

func GetContent(path string) ([]byte, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	return io.ReadAll(f)
}
