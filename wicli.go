package main

import (
	"config"
	"fmt"
	"os"
	prc "processing"
	ret "retriever"
	"strings"
)

const (
	RET_OK      = 0
	RET_USAGE   = 1
	RET_MISSING = 2
)

const (
	REDIRECT_MARKER = "#REDIRECT "
)

func main() {
	if len(os.Args) < 2 {
		config.PrintUsage()
		os.Exit(RET_USAGE)
	}

	cfg := config.GetConfig()

	query := os.Args[len(os.Args)-1]

	articleText := getArticleText(cfg, query)

	fmt.Printf("%s", articleText)
	os.Exit(RET_OK)
}

func getArticleText(cfg config.Config, query string) string {
	body := ret.Retrieve(cfg, query)

	//it is easy with raw option - just print it and we are done
	if cfg.Raw {
		fmt.Printf("%s", body)
		os.Exit(RET_OK)
	}

	//else post-process and check for missing articles, redirects etc
	articleMeta, articleText := prc.PostProcess(body, cfg)

	if articleMeta.Page.IsMissing {
		fmt.Fprintf(os.Stderr, "Could not find \"%s\"\n", query)
		os.Exit(RET_MISSING)
	}

	if cfg.FollowRedirects && strings.HasPrefix(articleText, REDIRECT_MARKER) {
		query := articleText[len(REDIRECT_MARKER):]
		fmt.Fprintf(os.Stderr, "Redirecting to \"%s\"\n", query)
		return getArticleText(cfg, query)
	}

	return articleText
}
