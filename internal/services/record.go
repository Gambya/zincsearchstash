package services

import (
	"encoding/json"
	"errors"

	"github.com/rs/zerolog/log"

	"zincsearchstash/pkg/zincsearch"
)

type Record interface {
	DataInsert(data interface{}) (err error)
}

type RecordService struct {
	zincSearch zincsearch.ZincSearchService
}

func NewRecordService(zincSearch zincsearch.ZincSearchService) *RecordService {
	return &RecordService{
		zincSearch: zincSearch,
	}
}

func (p *RecordService) DataInsert(data interface{}) (err error) {
	log.Trace().Msg("starting data insert")

	if data == nil {
		err = errors.New("error, message arrived null")
		log.Error().Err(err).Msg("error on get data")
		return
	}

	dataConverted, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("error on marshal data")
		return
	}

	_, err = p.zincSearch.Insert(string(dataConverted))
	if err != nil {
		log.Error().Err(err).Msg("error on insert data in zincsearch")
		return
	}

	return
}
