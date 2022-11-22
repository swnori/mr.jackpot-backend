package main

import (
	"log"

	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/rest"
	"mr.jackpot-backend/service/board"
	"mr.jackpot-backend/service/inventory"
	"mr.jackpot-backend/service/vui"
)

func main() {
	db.ConnectDB("127.0.0.1:3306", "mr.jackpot", "?parseTime=true")
	board.Initialize()
	inventory.Initialize()
	vui.Initialize()
 	log.Fatal(rest.RunAPI("0.0.0.0:8000"))
}