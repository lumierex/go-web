package amiorm

import (
	"ami-orm/log"
	"ami-orm/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
}

// orm实例化
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	// 心跳检测数据库连接
	if err := db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	e = &Engine{
		db: db,
	}
	log.Info("Connect database success")
	return e, nil

}

// 断开数据库连接
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

// 创建数据库操作session
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
