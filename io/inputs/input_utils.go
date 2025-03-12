package inputs

import (
	"bufio"
	"io"
	"iter"
)

// Lines returns a line iter.Seq for given Reader.
//
// Each line in the sequence is paired with a nil error.
// If an error is encountered, the final element of the sequence is an empty string paired with the error.
//
// This function strips trailing end-of-line marker of each line.
func Lines(r io.Reader) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		scanner := bufio.NewScanner(r)
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
