package session

import (
	"database/sql"
	"orm/dialect"
	"orm/log"
	"orm/schema"
	"strings"
)

type Session struct {
	db        *sql.DB
	refSchema *schema.Schema
	dialector dialect.Dialector
	sql       strings.Builder
	sqlVars   []interface{}
}

func New(db *sql.DB, dialector dialect.Dialector) *Session {
	return &Session{db: db, dialector: dialector}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(rawSql string, values ...interface{}) *Session {
	s.sql.WriteString(rawSql)
	s.sqlVars = values

	return s
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)

	if result, err = s.db.Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
