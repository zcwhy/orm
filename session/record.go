package session

import (
	"orm/clause"
	"reflect"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {
	if reflect.TypeOf(s.curModel) != reflect.TypeOf(values[0]) {
		s.Model(values[0])
	}

	s.clause.Set(clause.INSERT, s.refSchema.GetTableName(), s.refSchema.GetTableColumns())

	insertValues := make([]interface{}, 0)
	for _, value := range values {
		insertValues = append(insertValues, getValues(value))
	}
	s.clause.Set(clause.VALUES, insertValues...)

	sql, params := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, params...).Exec()
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func getValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))

	var fieldValues []interface{}
	for i := 0; i < destValue.Type().NumField(); i++ {
		fieldValues = append(fieldValues, destValue.Field(i).Interface())
	}

	return fieldValues
}
