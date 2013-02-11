package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	Api    = "http://%s.wikipedia.org/w/api.php?format=xml&action=query&titles=%s&prop=revisions&rvprop=content"
	Direct = "http://%s.wikipedia.org/wiki/%s"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Sprintf("Usage: %s article", os.Args[0]))
	}

	article := url.QueryEscape(os.Args[1])

	url := fmt.Sprintf(Api, "ru", article)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	type ApiResult struct {
		XMLName xml.Name
		Content string `xml:"query>pages>page>revisions>rev"`
	}

	v := ApiResult{}
	errXml := xml.Unmarshal(body, &v)
	if errXml != nil {
		log.Fatal(errXml)
	}

	fmt.Printf("%v", v.Content)
}

// func parseXml
