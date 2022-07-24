package zincsearch

import (
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type ZincSearchService interface {
	Insert(data string) (statusCode int, err error)
}

type ZincSearch struct {
	url  string
	user string
	pass string
}

func NewZincSearch(url, user, pass string) *ZincSearch {
	return &ZincSearch{
		url:  url,
		user: user,
		pass: pass,
	}
}

func (p *ZincSearch) Insert(data string) (statusCode int, err error) {
	req, err := http.NewRequest(http.MethodPost, p.url, strings.NewReader(data))
	if err != nil {
		log.Error().Err(err)
		return
	}
	req.SetBasicAuth(p.user, p.pass)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err)
		return
	}

	defer resp.Body.Close()
	statusCode = resp.StatusCode
	return
}
