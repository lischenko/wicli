package main

import (
	"config"
	"fmt"
	"os"
	prc "processing"
	ret "retriever"
)

const (
	RET_OK      = 0
	RET_USAGE   = 1
	RET_MISSING = 2
)

func main() {
	if len(os.Args) < 2 {
		config.PrintUsage()
		os.Exit(RET_USAGE)
	}

	cfg := config.GetConfig()

	query := os.Args[len(os.Args)-1]

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

	fmt.Printf("%s", articleText)
	os.Exit(RET_OK)
}
