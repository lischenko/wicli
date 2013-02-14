package main

import (
	"config"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"regexp"
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

	out := apiRes.Content

	if cfg.CleanupLinks {
		//clean-up named wikipedia links
		re := regexp.MustCompile("\\[\\[(.*?)\\|(.*?)\\]\\]")
		out = re.ReplaceAllString(out, "$2")

		//clean-up direct wikipedia links
		re = regexp.MustCompile("\\[\\[(.*?)\\]\\]")
		out = re.ReplaceAllString(out, "$1")
	}
	fmt.Printf("%v", out)
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
