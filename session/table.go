package session

import (
	"fmt"
	"orm/schema"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	s.refSchema = schema.Parse(value, s.dialector)
	s.curModel = value
	return s
}

func (s *Session) CreateTable() error {
	var columns []string
	for _, field := range s.refSchema.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}

	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", s.refSchema.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.refSchema.Name)).Exec()
	return err
}

func (s *Session) HasTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.refSchema.Name)).Exec()
	return err
}
