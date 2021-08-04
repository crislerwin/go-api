package main

import (
	"github.com/crislerwin/go-api/database"
	"github.com/crislerwin/go-api/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
