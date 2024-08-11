package strings2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "", Format("", 123))
	assert.Equal(t, "123", Format("{}", 123))
	assert.Equal(t, "123", Format("{0}", 123))
	assert.Equal(t, "       123", Format("{0:>10}", 123))
	assert.Equal(t, "123test", Format("{0}{1}", 123, "test"))
	assert.Equal(t, "-test-123", Format("-{1}-{0}", 123, "test"))

	assert.Equal(t, "[123]", Format("{}", []int{123}))
}

func TestFormatNamed(t *testing.T) {
	values := map[string]any{
		"name": "jack",
		"age":  16,
	}
	assert.Equal(t, "jack, 16", FormatNamed("{name}, {age}", values))
}

func TestFormatNamed2(t *testing.T) {
	type inner struct {
		male bool
	}

	values := struct {
		inner
		Name string
		Age  int
	}{
		Name: "jack",
		Age:  16,
	}
	assert.Equal(t, "jack, 16, false", FormatNamed2("{Name}, {Age}, {male}", values))
	assert.Equal(t, "jack, 16, false", FormatNamed2("{Name}, {Age}, {male}", &values))
}

func Test_formatter_parsePattern(t *testing.T) {
	cases := []struct {
		pattern string
		want    []segment
		wantErr bool
	}{
		{
			pattern: "{0}",
			want: []segment{
				{kind: kindArgument, value: "0"},
			},
		},
		{
			pattern: "{}",
			want: []segment{
				{kind: kindArgument, value: ""},
			},
		},
		{
			pattern: "0",
			want: []segment{
				{value: "0"},
			},
		},
		{
			pattern: "te{0:+.23}st",
			want: []segment{
				{value: "te"},
				{kind: kindArgument, value: "0", format: "+.23"},
				{value: "st"},
			},
		},
		{
			pattern: "te{name:5>}st",
			want: []segment{
				{value: "te"},
				{kind: kindArgument, value: "name", format: "5>"},
				{value: "st"},
			},
		},
	}
	for _, c := range cases {
		f := formatter{pattern: c.pattern}
		segments, err := f.parsePattern()
		if c.wantErr {
			assert.Error(t, err, "pattern '"+c.pattern+"' failed")
		} else {
			assert.NoError(t, err, "pattern '"+c.pattern+"' failed")
			assert.Equal(t, c.want, segments, "pattern '"+c.pattern+"' failed")
		}

	}
}

func Test_formatter_writeArgument(t *testing.T) {
	cases := []struct {
		format  string
		value   any
		want    string
		wantErr bool
	}{
		{format: ">1", value: 1, want: "1"},
		{format: ">1", value: -1, want: "-1"},
		{format: ">10", value: 1, want: "         1"},
		{format: "0>10", value: 1, want: "0000000001"},
		{format: "0>10", value: -1, want: "00000000-1"},
		{format: "0<10", value: -1, want: "-100000000"},
		{format: "0^10", value: -1, want: "0000-10000"},
		{format: "0>+10", value: -1, want: "00000000-1"},
		{format: "0>+10", value: 1, want: "00000000+1"},
		{format: "0> 10", value: 1, want: "00000000 1"},
		{format: "0>+10", value: 0, want: "00000000+0"},
		{format: ">+010", value: 0, want: "00000000+0"},
		{format: "a>+010", value: 0, want: "aaaaaaaa+0"},
		{format: ">#010", value: 123123, want: "0000123123"},
		{format: "=#010b", value: 1, want: "0b00000001"},
		{format: "=#010o", value: 123123, want: "0o00360363"},
		{format: ">#010x", value: 123123, want: "0000x1e0f3"},
		{format: "=#010x", value: 123123, want: "0x0001e0f3"},
		{format: "=#010x", value: -123123, want: "-0x001e0f3"},
		{format: "=#010X", value: -123123, want: "-0X001E0F3"},
		{format: "=#010x", value: uint(0xFFFFFFFF), want: "0xffffffff"},
		{format: "", value: 3.1415926, want: "3.1415926"},
		{format: ">.2f", value: 3.1415926, want: "3.14"},
		{format: "=#>.2f", value: 3.1415926, wantErr: true},
		{format: "=010.2f", value: 3.1415926, want: "0000003.14"},
		{format: "0:.2f", value: 3, want: "3.00"},
	}

	for _, c := range cases {
		f := formatter{}
		var sb strings.Builder
		err := f.writeArgument(&sb, c.value, c.format)
		if c.wantErr {
			assert.Error(t, err, fmt.Sprintf("format '%s', value '%v' failed", c.format, c.value))
		} else {
			assert.NoError(t, err, fmt.Sprintf("format '%s', value '%v' failed", c.format, c.value))
			assert.Equal(t, c.want, sb.String())
		}
	}
}
