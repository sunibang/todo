package main

import (
	"github.com/rs/zerolog/log"
	"github.com/sunibang/todo/api"
)

func main() {
	runGinServer()
}

func runGinServer() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start("localhost:9090")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
