package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/mohaila/echo-notes-rest/config"
	"github.com/mohaila/echo-notes-rest/server"
	"github.com/mohaila/echo-notes-rest/service"
	"github.com/mohaila/echo-notes-rest/store"
)

func main() {
	dc := config.GetDBConfig()

	store, err := store.NewStore(dc)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()
	log.Print("Connected to atabase: success")

	ns := service.NewNoteService(store)
	log.Print("Service created: success")

	sc := config.GetServerConfig()
	server := server.NewServer(sc, ns)

	log.Fatal(server.Start())
}
