package session

import (
	"amiorm/dialect"
	"amiorm/log"
	"amiorm/schema"
	"database/sql"
	"strings"
)

type Session struct {
	// sql.Open() 返回的句柄
	db *sql.DB

	dialect  dialect.Dialect
	refTable *schema.Schema

	sql     strings.Builder
	sqlVars []interface{}
}

// 创建Session
func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

// 清除session
func (s *Session) Clear() {
	//  清空strings.Builder
	s.sql.Reset()
	s.sqlVars = nil
}

// 获取db句柄
func (s *Session) DB() *sql.DB {
	return s.db
}

// 执行原生sql参数
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	result, err = s.DB().Exec(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

// 查询一条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows(row *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if row, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
