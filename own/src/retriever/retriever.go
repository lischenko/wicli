package retriever

import (
	"config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Retrieve(conf config.Config, article string) []byte {
	articleEscaped := url.QueryEscape(article)
	url := fmt.Sprintf(conf.UrlTemplate, conf.Lang, articleEscaped)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return body

}
