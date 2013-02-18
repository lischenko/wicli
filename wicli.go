package main

import (
	"config"
	"fmt"
	"os"
	prc "processing"
	ret "retriever"
)

func main() {
	if len(os.Args) < 2 {
		config.PrintUsage()
		os.Exit(0)
	}

	cfg := config.GetConfig()

	query := os.Args[len(os.Args)-1]

	article := ret.Retrieve(cfg, query)

	if !cfg.Raw {
		article = prc.PostProcess(article, cfg)
	}

	fmt.Printf("%s", article)
}
