package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mixanemca/example-gorilla-rest-api/config"
	"github.com/mixanemca/example-gorilla-rest-api/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Init config
	config.InitConfig()

	// Connect to DB
	if err := db.InitDatabase(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Conn.Close(context.Background())

	r := mux.NewRouter()
	http.Handle("/", r)

	addr := fmt.Sprintf("%s:%s", config.Cfg.App.ListenAddr, config.Cfg.App.Port)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
}
