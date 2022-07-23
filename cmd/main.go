package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"zincsearchstash/internal/broker"
	"zincsearchstash/internal/services"
	"zincsearchstash/internal/setup"
	"zincsearchstash/pkg/zincascii"
	"zincsearchstash/pkg/zincsearch"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := setup.Get()

	rmq := broker.NewRMQClient(cfg, cfg.BrokerUrl)

	zincSearch := zincsearch.NewZincSearch(fmt.Sprintf(cfg.ZincSearchUrl, cfg.ZincIndex), cfg.ZincUser, cfg.ZincPass)

	recordService := services.NewRecordService(zincSearch)

	setupGlobal(cfg.LogLevel)
	zincascii.Apply()

	err := rmq.RecordListening(context.Background(), recordService.DataInsert)
	if err != nil {
		panic(err)
	}
}

// SetupGlobal -
func setupGlobal(level int) {
	logWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	zerolog.SetGlobalLevel(zerolog.Level(level))
	log.Logger = zerolog.
		New(logWriter).
		With().
		Timestamp().
		Logger()

	log.Debug().Msg("Global logger configured successfuly")
}
