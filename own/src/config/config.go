package config

const (
	Api    = "http://%s.wikipedia.org/w/api.php?format=xml&action=query&titles=%s&prop=revisions&rvprop=content"
	Direct = "http://%s.wikipedia.org/wiki/%s"
)

type Config struct {
	UrlTemplate   string
	Lang          string
	StripBrackets bool
}

func PrintUsage() {
	print("Usage:\n\twicli article\n")
}

func getDefaultConfig() Config {
	return Config{UrlTemplate: Api, Lang: "ru", StripBrackets: true}
}

func GetConfig(args []string) Config {
	return getDefaultConfig() //for now
}
