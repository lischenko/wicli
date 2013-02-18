package processing

import (
	"config"
	"encoding/xml"
	"log"
	"regexp"
)

type Article struct {
	XMLName xml.Name
	Content string `xml:"query>pages>page>revisions>rev"`
}

func PostProcess(content string, cfg config.Config) string {
	parsed := parseXml(content)

	out := parsed.Content

	if cfg.CleanupLinks {
		//clean-up named wikipedia links
		re := regexp.MustCompile("\\[\\[(.*?)\\|(.*?)\\]\\]")
		out = re.ReplaceAllString(out, "$2")

		//clean-up direct wikipedia links
		re = regexp.MustCompile("\\[\\[(.*?)\\]\\]")
		out = re.ReplaceAllString(out, "$1")
	}

	if cfg.StripCategories {
		re := regexp.MustCompile("\\{\\{(.*?)\\}\\}\\n?")
		out = re.ReplaceAllString(out, "")
	}

	return out
}

func parseXml(xmlStr string) Article {
	a := Article{}

	errXml := xml.Unmarshal([]byte(xmlStr), &a)
	if errXml != nil {
		log.Fatal(errXml)
	}

	return a
}
