package schema

import (
	"orm/dialect"
	"reflect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  uint
}

func TestParse(t *testing.T) {
	TestDial := &dialect.SqlliteDialect{}

	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	// if schema.GetField("Name").Tag != "PRIMARY KEY" {
	// 	t.Fatal("failed to parse primary key")
	// }
}

func TestReflect(t *testing.T) {
	var num int
	t.Log(reflect.Indirect(reflect.ValueOf(num)).Type())

	// will panic
	// t.Log(reflect.Indirect(reflect.ValueOf(nil)).Type())

	t.Log(reflect.Indirect(reflect.ValueOf(&num)).Type())

	var nump *int
	val := reflect.New(reflect.TypeOf(nump))
	t.Log(val.Type())
}
