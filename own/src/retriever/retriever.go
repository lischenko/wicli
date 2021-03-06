package retriever

import (
	"config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	api = "http://%s.wikipedia.org/w/api.php?format=xml&action=query&titles=%s&prop=revisions&rvprop=content"
)

func Retrieve(cfg config.Config, query string) string {
	queryEscaped := url.QueryEscape(query)

	url := fmt.Sprintf(api, cfg.Lang, queryEscaped)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
