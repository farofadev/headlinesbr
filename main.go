package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/farofadev/primeirapagina/data"
	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {

	FetchHeadlines(data.Portals)
	//writeJSON(portals)
	//log.Println(data.Portals)

}

func writeJSON(data []Fact) {

	if len(data) == 0 {
		log.Println("No data to write")
		return
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	err = ioutil.WriteFile("rhinofacts"+fmt.Sprint(time.Now().Unix())+".json", file, 0644)

	if err != nil {
		log.Println("Unable to create json file", err)
	}

}

func FetchHeadlines(portals []data.Portal) []data.Portal {

	wg := sync.WaitGroup{}

	for index := range portals {
		wg.Add(1)
		go func(i int) {
			portal := &portals[i]
			collector := colly.NewCollector()

			collector.SetRequestTimeout(30 * time.Second)

			collector.OnHTML(portal.HeadlineSelector, func(element *colly.HTMLElement) {
				portal.Headline = element.Text
				if portal.HeadlineSelector == "" {
					log.Println("Retorno vazio da headline")
				}
			})

			collector.Visit(portal.Url)
			wg.Done()

		}(index)

	}
	wg.Wait()
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(portals)
	return portals

}
