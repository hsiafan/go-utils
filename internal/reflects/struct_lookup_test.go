package reflects

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructLookup_GetField(t *testing.T) {

	v := reflect.ValueOf(mockTestStruct())
	lookup := DefaultStructLookup()
	name, ok := lookup.Field(v, "Name")
	assert.True(t, ok)
	assert.Equal(t, "john", name.String())

	age, ok := lookup.Field(v, "age")
	assert.True(t, ok)
	assert.Equal(t, 18, int(age.Int()))

	male, ok := lookup.Field(v, "male")
	assert.True(t, ok)
	assert.Equal(t, true, male.Bool())

	weight, ok := lookup.Field(v, "weight")
	assert.True(t, ok)
	assert.Equal(t, 10.0, weight.Float())

	rich, ok := lookup.Field(v, "rich")
	assert.True(t, ok)
	assert.Equal(t, true, rich.Bool())
}

func mockTestStruct() any {
	return &testStruct{
		Name: "john",
		age:  18,
		testEmbedStruct2: &testEmbedStruct2{
			weight: 10,
			testEmbedStruct21: testEmbedStruct21{
				weight: 11,
				rich:   true,
			},
		},
		testEmbedStruct: testEmbedStruct{
			male: true,
		},
	}
}

type testStruct struct {
	testEmbedStruct
	*testEmbedStruct2
	Name string
	age  int
}

type testEmbedStruct struct {
	male bool
}

type testEmbedStruct2 struct {
	weight float32
	testEmbedStruct21
}

type testEmbedStruct21 struct {
	weight float32
	rich   bool
}
