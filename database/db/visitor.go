package db

import (
	"context"

	"mr.jackpot-backend/database/orm"
)


type VisitorLayer interface {
	CreateVisitor(identifier string) (int, error)
	GetVisitorByIdentifier(identifier string) (int, error)
}

type VisitorDB struct {
	DBAccessor
}

func NewVisitorDB() *VisitorDB {
	db := &VisitorDB{}
	db.q = NewAccessor()

	return db
}

func (db *VisitorDB) CreateVisitor(identifier string)  (int, error) {
	ctx := context.Background()

	result, err := db.q.CreateUser(ctx)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), db.q.CreateVisitor(ctx, orm.CreateVisitorParams{
		VisitorID: id,
		Identifier: identifier,
	})
}

func (db *VisitorDB) GetVisitorByIdentifier(identifier string) (int, error) {
	ctx := context.Background()

	id, err := db.q.GetVisitorID(ctx, identifier)
	return int(id), err
}