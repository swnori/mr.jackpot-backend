package main

import (
	"log"
	"mr.jackpot-backend/rest"
)

func main() {
	log.Fatal(rest.RunAPI("0.0.0.0:8000"))
}