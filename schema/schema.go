package schema

import (
	"go/ast"
	"orm/dialect"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Name   string
	Fields []*Field
}

// 把model解析为dialect对应的数据库schema
func Parse(model interface{}, d dialect.Dialector) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(model)).Type()
	s := &Schema{
		Name: modelType.Name(),
	}

	for i := 0; i < modelType.NumField(); i++ {
		fieldI := modelType.Field(i)
		if !fieldI.Anonymous && ast.IsExported(fieldI.Name) {
			f := &Field{
				Name: fieldI.Name,
				Type: d.DataTypeOf(fieldI.Type),
			}

			if v := fieldI.Tag.Get("geeorm"); v != "" {
				f.Tag = v
			}

			s.Fields = append(s.Fields, f)
		}
	}
	return s
}

func (s *Schema) GetTableColumns() []string {
	columns := []string{}
	for _, f := range s.Fields {
		columns = append(columns, f.Name)
	}

	return columns
}

func (s *Schema) GetTableName() string {
	return s.Name
}
