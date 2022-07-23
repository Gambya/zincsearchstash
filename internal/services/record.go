package services

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"zincsearchstash/internal/setup"
	"zincsearchstash/pkg/zincsearch"
)

type Record interface {
	DataInsert(data interface{}) (err error)
}

type RecordService struct {
	cfg        *setup.Config
	zincSearch zincsearch.ZincSearchService
}

func NewRecordService(zincSearch zincsearch.ZincSearchService) *RecordService {
	return &RecordService{
		zincSearch: zincSearch,
	}
}

func (p *RecordService) DataInsert(data interface{}) (err error) {
	log.Trace().Msg("starting data insert")

	dataConverted, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("error on marshal data")
		return
	}

	status, err := p.zincSearch.Insert(string(dataConverted))
	if err != nil {
		log.Error().Err(err).Msg("error on insert data in zincsearch")
		return
	}

	if status != http.StatusOK {
		log.Warn().Int("statuscode", status).Msg("warning on response")
		return
	}

	return
}
