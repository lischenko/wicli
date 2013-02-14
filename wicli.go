package main

import (
	"config"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"retriever"
)

func main() {
	if len(os.Args) < 2 {
		config.PrintUsage()
		os.Exit(0)
	}

	cfg := config.GetConfig()

	article := os.Args[len(os.Args)-1]

	body := retriever.Retrieve(cfg, article)

	apiRes := parseXml(body)
	fmt.Printf("%v", apiRes.Content)
}

type Article struct {
	XMLName xml.Name
	Content string `xml:"query>pages>page>revisions>rev"`
}

func parseXml(xmlBytes []byte) Article {
	a := Article{}

	errXml := xml.Unmarshal(xmlBytes, &a)
	if errXml != nil {
		log.Fatal(errXml)
	}

	return a
}
