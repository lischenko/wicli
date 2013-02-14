package config

import "flag"

type Config struct {
	Lang       string
	StripLinks bool
}

var cfg Config

func init() {
	flag.StringVar(&cfg.Lang, "lang", "en", "Language of the Wikipedia to be queried")
	flag.BoolVar(&cfg.StripLinks, "stripLinks", true, "Wheter or not to strip links")

	flag.Parse()
}

func PrintUsage() {
	print("Usage:\n\twicli [flags] article\nFlags and their default values:\n")
	flag.PrintDefaults()
}

func GetConfig() Config {
	return cfg
}
