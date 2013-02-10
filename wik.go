package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
)

const (
	Api    = "http://en.wikipedia.org/w/api.php?format=xml&action=query&titles=%s&prop=revisions&rvprop=content"
	Direct = "http://en.wikipedia.org/wiki/%s"
)

func main() {
	// if len(os.Args) != 2 {
	// 	log.Fatal(fmt.Sprintf("Usage: %s article", os.Args[0]))
	// }

	// article := os.Args[1]

	article := "Experiment"
	url := fmt.Sprintf(Api, article)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
}
