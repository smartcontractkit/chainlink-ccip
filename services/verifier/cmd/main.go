package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	lvlStr := os.Getenv("VERIFIER_LOG_LEVEL")
	if lvlStr == "" {
		lvlStr = "info"
	}
	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		panic(err)
	}
	l := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(lvl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Verfifier is running!\n")
	})

	l.Info().Msgf("Verifier is running on port %s", ":8100")
	log.Fatal().Err(http.ListenAndServe(":8100", nil)).Send()
}
