package main

import (
	"log"

	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/rest"
	"mr.jackpot-backend/service/board"
	"mr.jackpot-backend/service/stock"
	"mr.jackpot-backend/service/vui"
)

func main() {
	if err := db.ConnectDB("127.0.0.1:3306", "mr.jackpot", "?parseTime=true"); err != nil {
		panic(err)
	}
	if err := board.Initialize(); err != nil {
		panic(err)
	}
	if err := stock.Initialize(); err != nil {
		panic(err)
	}
	if err := vui.Initialize(); err != nil {
		panic(err)
	}

	log.Fatal(rest.RunAPI("0.0.0.0:8000"))
}

