package jsons

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/hsiafan/go-utils/strings2"
)

// Marshal encode value as json bytes, with optional MarshalOption.
// Unlike the json.Marshal, EscapeHTML is disabled by default.
func Marshal(w io.Writer, v any, options ...MarshalOption) ([]byte, error) {
	var b bytes.Buffer
	if err := MarshalTo(&b, v, options...); err != nil {
		return nil, err
	}
	bs := b.Bytes()
	bs = bs[:len(bs)-1]
	return bs, nil
}

// MarshalString encode value as json string, with optional MarshalOption.
// Unlike the json.Marshal, EscapeHTML is disabled by default.
func MarshalString(w io.Writer, v any, options ...MarshalOption) (string, error) {
	var sb strings.Builder
	if err := MarshalTo(&sb, v, options...); err != nil {
		return "", err
	}
	s := sb.String()
	s = s[:len(s)-1]
	return s, nil
}

// MarshalTo encode value as json to a io.Writer, with optional MarshalOption.
// Unlike the json.Marshal, EscapeHTML is disabled by default.
// This func write a trailing empty line at the end of json value.
func MarshalTo(w io.Writer, v any, options ...MarshalOption) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	for _, option := range options {
		option(enc)
	}

	return enc.Encode(v)
}

// MarshalOption is a func that sets json.Encoder options.
type MarshalOption func(*json.Encoder)

// EscapeHTML is a json MarshalOption enable html escapes.
func EscapeHTML() MarshalOption {
	return func(e *json.Encoder) {
		e.SetEscapeHTML(true)
	}
}

// IndentWith is a json MarshalOption that sets indent.
func IndentWith(prefix, indent string) MarshalOption {
	return func(e *json.Encoder) {
		e.SetIndent(prefix, indent)
	}
}

// UnmarshalString is like json.Unmarshal, but takes a string as input.
func UnmarshalString(s string, v any) error {
	return json.Unmarshal(strings2.AsBytes(s), v)
}
