// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package orm

import (
	"database/sql"
	"time"
)

type BoardEntity struct {
	EntityID int32
	TargetID int32
	Name     string
	Tag      string
	Price    int32
}

type CouponIssued struct {
	CouponID    int64
	Code        string
	Amount      int32
	Title       sql.NullString
	Description sql.NullString
	CreatedAt   time.Time
	ExpiresAt   time.Time
}

type CouponOwned struct {
	CouponID int64
	OwnerID  int64
	Valid    bool
}

type Customer struct {
	CustomerID int64
	Status     bool
	Name       string
	Address    sql.NullString
	Phone      sql.NullString
	Orders     int32
	Rating     int32
	Paid       int32
}

type CustomerAuth struct {
	ID         string
	Password   string
	CustomerID int64
}

type DeliveryInfo struct {
	OrderID int64
	Name    string
	Address string
	Phone   string
	Message sql.NullString
}

type Dinner struct {
	DinnerID int32
	EntityID int32
}

type DinnersMenu struct {
	DinnerID     int32
	MenuID       int32
	DefaultCount int32
}

type DinnersStyle struct {
	DinnerID int32
	StyleID  int32
}

type EntityCount struct {
	CountID  int32
	Count    int32
	TargetID int32
}

type EntityType struct {
	TypeID   int32
	Typename string
}

type Menu struct {
	MenuID      int32
	EntityID    int32
	TypeID      int32
	Option1Name sql.NullString
	Option2Name sql.NullString
}

type MenuOption1 struct {
	MenuID   int32
	OptionID int32
	EntityID int32
}

type MenuOption2 struct {
	MenuID   int32
	OptionID int32
	EntityID int32
}

type MenuRole struct {
	MenuID int32
	RoleID int32
}

type MenuType struct {
	ID   int32
	Name string
}

type Order struct {
	OrderID      int64
	UserID       int64
	Price        int32
	Deposit      int32
	Discount     int32
	ReservatedAt time.Time
}

type OrderState struct {
	OrderID int64
	StateID int32
}

type OrderedDinner struct {
	ID      int64
	OrderID int64
	StyleID int32
	Amount  int32
}

type OrderedMenu struct {
	ID         int64
	OrderID    int64
	DinnerID   int64
	MenutypeID int32
	MenuID     int32
	Option1ID  int32
	Option2ID  int32
	Count      int32
	Price      sql.NullInt32
}

type PreOrderChoice struct {
	SeqID   int32
	Tag     string
	Message string
}

type PreOrderChoiceNxtSeq struct {
	SeqID int32
	NxtID int32
}

type ProOrderChoice struct {
	SeqID   int32
	Tag     string
	Target  string
	Message string
	TypeID  int32
}

type ProOrderChoiceNxtSeq struct {
	SeqID int32
	NxtID int32
}

type Role struct {
	RoleID int32
	Tag    string
	Name   string
}

type Staff struct {
	StaffID int64
	Status  bool
	RoleID  int32
	Name    string
	Score   int32
}

type StaffAuth struct {
	Code    string
	StaffID int64
}

type State struct {
	StateID int32
	Name    string
}

type Stock struct {
	StockID int64
	Name    string
	Count   int32
}

type Style struct {
	StyleID     int32
	EntityID    int32
	Description string
}

type User struct {
	UserID int64
}

type Visitor struct {
	Identifier string
	VisitorID  int64
}
