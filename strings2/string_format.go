package strings2

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/hsiafan/go-utils/internal/reflects"
)

// Format format a string use python-PEP3101 style, with positional arguments.
//
// # Format specifier:
//
//	[[fill]align][sign][#][0][minimumwidth][.precision][type]
//
// The optional align feature can be one of the following:
//   - '<': left-aligned
//   - '>': right-aligned (this is the default)
//   - '=': padding after the sign, but before the digits (e.g., for +000042)
//   - '^': centered
//
// If an align flag is defined, a 'fill' character can also be defined. If undefined, space (' ') will be used.
//
// The optional 'sign' is only valid for numeric types and can be:
//   - '+': Show sign for both positive and negative numbers
//   - '-': Show sign only for negative numbers (default)
//   - ' ': use a leading space for positive numbers
//
// If # is present, when using the binary, octal, or hex types, a '0b', '0o', or '0x' will be prepended, respectively.
//
// The minimumwidth field specifies a minimum width, which is helpful when used with alignment. If preceded with a zero, numbers will be zero-padded.
//
// The precision field specifies a precision width for float types(and also complex types).
//
// The 'type' format determines what type the value will be formatted as. For integers:
//   - 'b' - Binary, base 2
//   - 'd' - Decimal, base 10 (default)
//   - 'o' - Octal, base 8
//   - 'x' - Hexadecimal, base 16
//   - 'X' - Hexadecimal, base 16, using upper-case letters
//
// For floats and complex numbers:
//   - 'e' - Scientific notation
//   - 'E' - Similar to e, but uppercase
//   - 'f' - Fixed point, displays the number as a fixed-point number.
//   - 'g' - General format, prints as a fixed point unless it's too large, then switches to scientific notation. (default)
//   - 'G' - Similar to g, but uses capital letters
func Format(pattern string, values ...any) string {
	f := formatter{pattern: pattern}
	s, err := f.format(values...)
	if err != nil {
		panic(err)
	}
	return s
}

// FormatNamed format a string use python-PEP3101 style, with name-value arguments.
//
// For the format specifier, see [Format].
func FormatNamed[V any](pattern string, values map[string]V) string {
	f := formatter{pattern: pattern}
	s, err := f.formatNamed(func(name string) (any, bool) {
		v, ok := values[name]
		return v, ok
	})
	if err != nil {
		panic(err)
	}
	return s
}

// FormatNamed format a string use python-PEP3101 style, with name-value arguments.
// The values must be a struct, or a pointer to struct.
//
// For the format specifier, see [Format].
func FormatNamed2(pattern string, values any) string {
	rv := reflect.ValueOf(values)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		panic(errors.New("values must be map[string]any or struct"))
	}

	f := formatter{pattern: pattern}
	lookup := reflects.DefaultStructLookup()
	s, err := f.formatNamed(func(name string) (any, bool) {
		v, ok := lookup.Field(rv, name)
		if ok {
			return v.Interface(), true
		} else {
			return nil, false
		}
	})
	if err != nil {
		panic(err)
	}
	return s
}

type formatter struct {
	pattern string
}

type kind int

const (
	kindText kind = iota
	kindArgument
)

type segment struct {
	kind   kind
	value  string
	format string // only for argument
}

type state_ int

const (
	stateText state_ = iota
	stateOpenBrace
	stateCloseBrace
	stateArgument
)

func (f *formatter) format(values ...any) (string, error) {
	segments, err := f.parsePattern()
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	for i, seg := range segments {
		switch seg.kind {
		case kindText:
			sb.WriteString(seg.value)
		case kindArgument:
			var index int
			if len(seg.value) == 0 {
				index = i
			} else {
				index, err = strconv.Atoi(seg.value)
				if err != nil {
					return "", fmt.Errorf("argument index '%v' is not a number", seg.value)
				}
			}
			if index < 0 || index > len(values) {
				return "", fmt.Errorf("argument index %v out of range", index)
			}
			f.writeArgument(&sb, values[index], seg.format)
		}
	}
	return sb.String(), nil
}

func (f *formatter) formatNamed(lookup func(string) (any, bool)) (string, error) {
	segments, err := f.parsePattern()
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	for _, seg := range segments {
		switch seg.kind {
		case kindText:
			sb.WriteString(seg.value)
		case kindArgument:
			if len(seg.value) == 0 {
				return "", errors.New("argument name cannot be empty")
			}
			v, ok := lookup(seg.value)
			if !ok {
				return "", fmt.Errorf("argument with name '%s' not exists", seg.value)
			}
			f.writeArgument(&sb, v, seg.format)
		}
	}
	return sb.String(), nil
}

func (f *formatter) parsePattern() ([]segment, error) {
	var segments []segment
	var sb strings.Builder
	var state state_
	for _, r := range f.pattern {
		switch r {
		case '{':
			switch state {
			case stateText:
				state = stateOpenBrace
			case stateOpenBrace:
				sb.WriteByte('{')
			default:
				return nil, errors.New("unexpected char '{'")
			}
		case '}':
			switch state {
			case stateText:
				state = stateCloseBrace
			case stateOpenBrace:
				// {}
				segments = append(segments, segment{
					kind: kindArgument,
				})
				state = stateText
			case stateCloseBrace:
				sb.WriteByte('}')
			case stateArgument:
				name, format, _ := strings.Cut(strings.Clone(sb.String()), ":")
				segments = append(segments, segment{
					kind:   kindArgument,
					value:  name,
					format: format,
				})
				state = stateText
				sb.Reset()
			default:
				return nil, errors.New("unexpected char '{'")
			}
		default:
			switch state {
			case stateText:
				sb.WriteRune(r)
			case stateOpenBrace:
				str := strings.Clone(sb.String())
				if len(str) != 0 {
					segments = append(segments, segment{
						kind:  kindText,
						value: strings.Clone(sb.String()),
					})
				}
				sb.Reset()
				sb.WriteRune(r)
				state = stateArgument
				continue
			case stateArgument:
				sb.WriteRune(r)
			default:
				return nil, fmt.Errorf("unexpected char '%v'", r)

			}
		}
	}
	switch state {
	case stateText:
		str := strings.Clone(sb.String())
		if len(str) != 0 {
			segments = append(segments, segment{kind: kindText, value: str})
		}
	default:
		return nil, errors.New("unexpected end")
	}
	return segments, nil
}

type formatStep int

// [[fill]align][sign][#][0][minimumwidth][.precision][type]
const (
	_ formatStep = iota
	formatAlignFill
	formatAlign
	formatSign
	formatSharp
	formatZero
	formatWidth
	formatPrecision
	formatType
)

func (f *formatter) writeArgument(sb *strings.Builder, v any, format string) error {
	// for all
	var fill rune
	var align byte
	var minWidth = 0
	// for all numbers
	var sign byte
	var pad rune
	var _type byte
	// for int values
	intBase := 10
	upperCase := false // only for hex
	prependPrefix := false
	// for floats and complex
	var floatFormat byte = 'f'
	floatPrec := -1

	// parse format
	var lastStep formatStep
	// read first chat to see if is fill
	if len(format) > 0 {
		first, firstSize := utf8.DecodeRuneInString(format)
		if firstSize < len(format) {
			second, _ := utf8.DecodeRuneInString(format[firstSize:])
			switch second {
			case '>', '<', '=', '^':
				fill = first
				format = format[firstSize:]
				lastStep = formatAlignFill
			}
		}
	}

	for _, r := range format {
		switch r {
		case '>', '<', '=', '^':
			if lastStep >= formatAlign {
				return errors.New("invalid format: " + format)
			}
			align = byte(r)
			lastStep = formatAlign
		case '+', '-', ' ':
			if lastStep >= formatSign {
				return errors.New("invalid format: " + format)
			}
			sign = byte(r)
			lastStep = formatSign
		case '#':
			if lastStep >= formatSharp {
				return errors.New("invalid format: " + format)
			}
			prependPrefix = true
			lastStep = formatSharp
		case '0':
			if lastStep < formatPrecision {
				if lastStep < formatWidth {
					pad = '0'
					lastStep = formatWidth
				}
			}
			fallthrough
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if lastStep < formatPrecision {
				if lastStep > formatWidth {
					return errors.New("invalid format: " + format)
				}
				minWidth = minWidth*10 + int(r-'0')
				lastStep = formatWidth
			} else {
				if lastStep > formatPrecision {
					return errors.New("invalid format: " + format)
				}
				if floatPrec == -1 {
					floatPrec = int(r - '0')
				} else {
					floatPrec = floatPrec*10 + int(r-'0')
				}
				lastStep = formatPrecision
			}
		case '.':
			if lastStep >= formatPrecision {
				return errors.New("invalid format: " + format)
			}
			lastStep = formatPrecision

		case 'b', 'd', 'o', 'x', 'X':
			if lastStep >= formatType {
				return errors.New("invalid format: " + format)
			}
			_type = 'i'
			switch r {
			case 'd':
				intBase = 10
			case 'b':
				intBase = 2
			case 'o':
				intBase = 8
			case 'X':
				upperCase = true
				fallthrough
			case 'x':
				intBase = 16
			}
			lastStep = formatType
		case 'e', 'E', 'f', 'g', 'G':
			if lastStep >= formatType {
				return errors.New("invalid format: " + format)
			}
			_type = 'f'
			floatFormat = byte(r)
			lastStep = formatType
		}
	}

	// check type
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:

	case float32, float64, complex64, complex128:
		if _type == 'i' {
			return errors.New("float value cannot set int type")
		}
	default:
		if _type != 0 {
			return errors.New("non number type cannot set number type")
		}
	}

	var isInt = false
	var isFloat = false
	var negative = false // for sign
	// value to string
	var s string
	switch vv := v.(type) {
	case bool:
		s = strconv.FormatBool(vv)
	case int, int8, int16, int32, int64:
		iv := f.toInt64(v)
		negative = iv < 0
		if _type != 'f' {
			s = strconv.FormatInt(int64(iv), intBase)
			isInt = true
		} else {
			s = strconv.FormatFloat(float64(iv), floatFormat, floatPrec, 64)
			isFloat = true
		}
	case uint, uint8, uint16, uint32, uint64:
		uiv := f.toUint64(v)
		if _type != 'f' {
			s = strconv.FormatUint(uint64(uiv), intBase)
			isInt = true
		} else {
			s = strconv.FormatFloat(float64(uiv), floatFormat, floatPrec, 64)
			isFloat = true
		}
	case float32:
		s = strconv.FormatFloat(float64(vv), floatFormat, floatPrec, 32)
		negative = vv < 0
		isFloat = true
	case float64:
		s = strconv.FormatFloat(vv, floatFormat, floatPrec, 64)
		negative = vv < 0
		isFloat = true
	case complex64:
		s = strconv.FormatComplex(complex128(vv), floatFormat, floatPrec, 64)
	case complex128:
		s = strconv.FormatComplex(vv, floatFormat, floatPrec, 128)
	case string:
		s = vv
	case fmt.Stringer:
		s = vv.String()
	default:
		s = fmt.Sprintf("%v", vv)
	}
	if upperCase {
		s = strings.ToUpper(s)
	}

	var prefix string
	if isInt && prependPrefix {
		switch intBase {
		case 2:
			prefix = "0b"
		case 8:
			prefix = "0o"
		case 16:
			if upperCase {
				prefix = "0X"
			} else {
				prefix = "0x"
			}

		}
	}

	var signStr string
	if isInt || isFloat {
		switch sign {
		case '+':
			if negative {
				signStr = "-"
				s = s[1:]
			} else {
				signStr = "+"
			}
		case ' ':
			if negative {
				signStr = "-"
				s = s[1:]
			} else {
				signStr = " "
			}
		case '-':
			fallthrough
		default:
			if negative {
				signStr = "-"
				s = s[1:]
			}
		}

	}
	var finalPad rune
	if fill != 0 {
		finalPad = fill
	} else if pad != 0 {
		finalPad = pad
	} else {
		finalPad = ' '
	}

	toAlign := minWidth - len(prefix) - len(s) - len(signStr)
	if toAlign > 0 {
		if align == '>' {
			for i := 0; i < toAlign; i++ {
				sb.WriteRune(finalPad)
			}
		} else if align == '^' {
			for i := 0; i < toAlign/2; i++ {
				sb.WriteRune(finalPad)
			}
		}
	}

	sb.WriteString(signStr)
	sb.WriteString(prefix)
	if toAlign > 0 && align == '=' {
		for i := 0; i < toAlign; i++ {
			sb.WriteRune(finalPad)
		}
	}

	sb.WriteString(s)

	if toAlign > 0 {
		if align == '<' {
			for i := 0; i < toAlign; i++ {
				sb.WriteRune(finalPad)
			}
		} else if align == '^' {
			for i := 0; i < toAlign-toAlign/2; i++ {
				sb.WriteRune(finalPad)
			}
		}
	}
	return nil
}

func (f *formatter) toInt64(v any) int64 {
	switch vv := v.(type) {
	case int:
		return int64(vv)
	case int8:
		return int64(vv)
	case int16:
		return int64(vv)
	case int32:
		return int64(vv)
	case int64:
		return vv
	default:
		panic("")
	}
}

func (f *formatter) toUint64(v any) uint64 {
	switch vv := v.(type) {
	case uint:
		return uint64(vv)
	case uint8:
		return uint64(vv)
	case uint16:
		return uint64(vv)
	case uint32:
		return uint64(vv)
	case uint64:
		return vv
	default:
		panic("")
	}
}
