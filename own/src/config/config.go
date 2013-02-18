package config

import "flag"

type Config struct {
	Lang            string
	Raw             bool
	CleanupLinks    bool
	StripCategories bool
	StripInterwiki  bool
}

var cfg Config

func init() {
	flag.StringVar(&cfg.Lang, "lang", "en", "Language of the Wikipedia to be queried")
	flag.BoolVar(&cfg.CleanupLinks, "cleanupLinks", true, "Whether or not to clean-up links")
	flag.BoolVar(&cfg.StripCategories, "stripCategories", true, "Whether or not to strip categories")
	flag.BoolVar(&cfg.Raw, "raw", false, "Supresses all post-processing")
	flag.BoolVar(&cfg.StripInterwiki, "stripInterwiki", true, "Whether or not to strip interwiki (language) links")

	flag.Parse()
}

func PrintUsage() {
	print("Usage:\n\twicli [flags] article\n\nFlags and their default values:\n")
	flag.PrintDefaults()
}

func GetConfig() Config {
	return cfg
}
