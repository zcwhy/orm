package dialect

import (
	"reflect"
	"testing"
)

func TestDateOfType(t *testing.T) {
	d := &SqlliteDialect{}

	tests := []struct {
		input    any
		expected string
	}{
		{true, "bool"},
		{int(0), "integer"},
		{int8(0), "integer"},
		{int16(0), "integer"},
		{int32(0), "integer"},
		{uint(0), "integer"},
		{uint8(0), "integer"},
		{uint16(0), "integer"},
		{uint32(0), "integer"},
		{uintptr(0), "integer"},
		{int64(0), "bigint"},
		{uint64(0), "bigint"},
		{"hello", "text"},
		{3.14, ""},           // float should not match
		{[]byte("data"), ""}, // not handled
	}

	for _, tt := range tests {
		t.Run(reflect.TypeOf(tt.input).String(), func(t *testing.T) {
			typ := reflect.TypeOf(tt.input)
			got := d.DataTypeOf(typ)
			if got != tt.expected {
				t.Errorf("DataTypeOf(%v) = %v; expected %v", typ, got, tt.expected)
			}
		})
	}
}
