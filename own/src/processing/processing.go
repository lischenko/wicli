package processing

import (
	"config"
	"encoding/xml"
	"log"
	"regexp"
)

const (
	missingMarker = "NO"
)

type PageMeta struct {
	Id           int    `xml:"pageid,attr"`
	Content      string `xml:"revisions>rev"`
	IsMissingRaw string `xml:"missing,attr"`
	IsMissing    bool   `xml:"-"`
}

type ArticleMeta struct {
	XMLName xml.Name
	// Content string `xml:"query>pages>page>revisions>rev"`
	Page PageMeta `xml:"query>pages>page"`
}

func PostProcess(content string, cfg config.Config) (ArticleMeta, string) {
	parsed := parseXml(content)

	out := parsed.Page.Content

	if cfg.StripInterwiki {
		out = stripInterwiki(out)
	}

	if cfg.StripCategories {
		re := regexp.MustCompile("\\[\\[Category:(.*?)\\]\\]\\n?")
		out = re.ReplaceAllString(out, "")
	}

	if cfg.CleanupLinks {
		//clean-up named wikipedia links
		re := regexp.MustCompile("\\[\\[([^\\]]*?)\\|([^\\[]*?)\\]\\]")
		out = re.ReplaceAllString(out, "$2")

		//clean-up direct wikipedia links
		re = regexp.MustCompile("\\[\\[(.*?)\\]\\]")
		out = re.ReplaceAllString(out, "$1")

		//clean-up named external links
		re = regexp.MustCompile("\\[(.*?) (.*?)\\]")
		out = re.ReplaceAllString(out, "$2")

		//clean-up external links
		re = regexp.MustCompile("\\[(.*?)\\]")
		out = re.ReplaceAllString(out, "$1")
	}

	// if cfg.StripMacros {
	// 	re := regexp.MustCompile("\\{\\{(.*?)\\}\\}\\n?")
	// 	out = re.ReplaceAllString(out, "")
	// }

	if cfg.StripComments {
		re := regexp.MustCompile("<!--(.*?)-->\\n?")
		out = re.ReplaceAllString(out, "")
	}

	return parsed, out
}

func parseXml(xmlStr string) ArticleMeta {
	p := PageMeta{IsMissingRaw: missingMarker}
	a := ArticleMeta{Page: p}

	errXml := xml.Unmarshal([]byte(xmlStr), &a)
	if errXml != nil {
		log.Fatal(errXml)
	}

	// interpret IsMissingRaw manually (couldn't do it declaratively with go xml unmarshaller)
	a.Page.IsMissing = !(a.Page.IsMissingRaw == missingMarker)

	return a
}
