package mysql

import (
	"comm/logger"
	"comm/micro"
	"errors"
	"time"

	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/log"

	_ "github.com/go-sql-driver/mysql"
)

var dbmap map[string]*xorm.Engine

func init() {
	dbmap = map[string]*xorm.Engine{}
}

// InitDB defined TODO
func InitDB(key string) (*xorm.Engine, error) {
	dbUri := micro.Key(key).String()
	if dbUri == "" {
		return nil, errors.New("not found uri")
	}

	if dbmap[dbUri] != nil {
		return dbmap[dbUri], nil
	}
	db, err := xorm.NewEngine("mysql", dbUri)
	if err != nil {
		return nil, err
	}

	logger := log.NewSimpleLogger(logger.Out())
	logger.ShowSQL(true)

	db.SetLogger(logger)
	db.SetConnMaxLifetime(time.Duration(10) * time.Minute)
	db.SetMaxIdleConns(15)
	db.SetMaxOpenConns(50)
	dbmap[dbUri] = db
	return db, nil
}
