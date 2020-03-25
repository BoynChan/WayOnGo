package gdorm

import (
	"7-days-gdorm/gdorm/dialect"
	"7-days-gdorm/gdorm/log"
	"7-days-gdorm/gdorm/session"
	"database/sql"
)

// Author:Boyn
// Date:2020/3/23

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	dialect, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error("can not get the corresponding dialect of this type of db")
	}
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping message to make sure that db has been connected
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db, dialect: dialect}
	log.Info("Connect Database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Fail to close db connect")
	}
	log.Info("Close db successfully")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
