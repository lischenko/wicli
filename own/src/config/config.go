package config

import "flag"

type Config struct {
	Lang            string
	Raw             bool
	CleanupLinks    bool
	StripCategories bool
}

var cfg Config

func init() {
	flag.StringVar(&cfg.Lang, "lang", "en", "Language of the Wikipedia to be queried")
	flag.BoolVar(&cfg.CleanupLinks, "cleanupLinks", true, "Wheter or not to clean-up links")
	flag.BoolVar(&cfg.StripCategories, "stripCategories", true, "Wheter or not to strip categories")
	flag.BoolVar(&cfg.Raw, "raw", false, "Supresses all post-processing")

	flag.Parse()
}

func PrintUsage() {
	print("Usage:\n\twicli [flags] article\n\nFlags and their default values:\n")
	flag.PrintDefaults()
}

func GetConfig() Config {
	return cfg
}
