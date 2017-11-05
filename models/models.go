package models

import (
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"github.com/Dell-/goci/pkg/settings"
	"path"
	"fmt"
	"database/sql"
)

// Engine represents a XORM engine or session.
type Engine interface {
	Delete(interface{}) (int64, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	Id(interface{}) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Iterate(interface{}, xorm.IterFunc) error
	Sql(string, ...interface{}) *xorm.Session
	Table(interface{}) *xorm.Session
	Where(interface{}, ...interface{}) *xorm.Session
}

var (
	tables []interface{}
	engine *xorm.Engine
)

func init() {
	tables = append(tables,
		new(User),
	)
}

func initLogger() {
	f, err := os.Create(path.Join(settings.APP.LogPath, "db.log"))
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))

	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
}

func newEngine() (*xorm.Engine, error) {
	var connStr string
	connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true",
		settings.DB.Username, settings.DB.Password, settings.DB.Host, settings.DB.Name, settings.DB.Charset)

	return xorm.NewEngine("mysql", connStr)
}

func NewEngine() (err error) {
	engine, err = newEngine()
	if err != nil {
		return fmt.Errorf("Fail to connect to database: %v", err)
	}

	engine.SetMapper(core.GonicMapper{})

	initLogger()

	if err = engine.StoreEngine("InnoDB").Sync2(tables...); err != nil {
		return fmt.Errorf("Sync database struct error: %v\n", err)
	}

	return nil
}

func Ping() error {
	return engine.Ping()
}
