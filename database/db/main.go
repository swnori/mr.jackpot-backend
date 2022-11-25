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



var DB *sql.DB

func NewAccessor() *ormpkg.Queries {
	return ormpkg.New(DB)
}

func ConnectDB(address, dbname, dsn string) (err error) {
	DB, err = sql.Open("mysql", fmt.Sprintf("root@tcp(%s)/%s%s", address, dbname, dsn))
	return err
}

