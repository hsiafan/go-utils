package files

import (
	"bufio"
	"iter"
	"os"
	"path/filepath"
)

// Create creates the named file if not exists.
// If file already exists, will truncate the file if truncate == true, append the file otherwise.
func CreateFile(path string, truncate bool) (*os.File, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0770); err != nil {
		return nil, err
	}
	flag := os.O_RDWR | os.O_CREATE
	if truncate {
		flag |= os.O_TRUNC
	} else {
		flag |= os.O_APPEND
	}
	return os.OpenFile(path, flag, 0666)
}

// Lines returns a line iter.Seq for given file.
//
// Each line in the sequence is paired with a nil error.
// If an error is encountered, the final element of the sequence is an empty string paired with the error.
//
// This function strips trailing end-of-line marker of each line.
func Lines(path string) iter.Seq2[string, error] {

	return func(yield func(string, error) bool) {
		f, err := os.Open(path)
		if err != nil {
			yield("", err)
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if !yield(scanner.Text(), nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield("", err)
		}
	}
}
