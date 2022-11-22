package db

import (
	"database/sql"
	"fmt"

	_ "gorm.io/driver/mysql"
	ormpkg "mr.jackpot-backend/database/orm"
)



type DBAccessor struct {
	q *ormpkg.Queries
}

var Accessor *ormpkg.Queries

func NewAccessor() *ormpkg.Queries {
	return Accessor
}


func ConnectDB(address, dbname, dsn string) error {

	db, err := sql.Open("mysql", fmt.Sprintf("root@tcp(%s)/%s%s", address, dbname, dsn))
	if err != nil {
		return err
	}
	Accessor = ormpkg.New(db)
	return nil
}

