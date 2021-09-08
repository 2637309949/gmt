package mysql

import (
	"comm/logger"
	"comm/micro"
	"errors"
	"time"

	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/log"
)

var dbmap map[string]*xorm.Engine

// InitDb defined TODO
func InitDb(key string) (*xorm.Engine, error) {
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

	db.SetConnMaxLifetime(time.Duration(10) * time.Minute)
	db.SetMaxIdleConns(15)
	db.SetMaxOpenConns(50)
	dbmap[dbUri] = db
	return db, nil
}
