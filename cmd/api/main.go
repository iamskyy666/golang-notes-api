package main

import (
	"fmt"
	"log"

	"github.com/callmeskyy111/notes-api/internal/config"
	"github.com/callmeskyy111/notes-api/internal/db"
	"github.com/callmeskyy111/notes-api/internal/server"
)

// config -> db -> router -> run server

func main() {
	cfg,err:=config.Load()
	if err != nil {
		log.Fatal("⚠️ Config Error!", err)
	}

	client,database,err:=db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("⚠️ DB Error!", err)
	}

	defer func ()  {
		if err:=db.DisconnectDB(client);err!=nil{
			log.Fatal("⚠️ MONGO Disconnect-Error!", err)
		}
	}()

	// create router
	router:=server.NewRouter(database)
	addr:=fmt.Sprintf(":%s",cfg.ServerPort)
	if err:=router.Run(addr);err!=nil{
			log.Fatalf("⚠️ Failed to connect to server: %v", err)
		}
}
