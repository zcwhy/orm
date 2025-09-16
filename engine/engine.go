package engine

import (
	"database/sql"
	"orm/dialect"
	"orm/log"
	"orm/session"
)

type Engine struct {
	db *sql.DB
	session.Session
}

func Open(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	dialector, ok := dialect.GetDialector(driver)
	if !ok {
		log.Error("Cloud not find diaector for:%s", driver)
		return
	}

	e = &Engine{db: db, Session: *session.New(db, dialector)}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}
