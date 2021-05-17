package scheduler

import (
	"log"
	"time"

	"github.com/farofadev/headlinesbr/data"
)

func Run() {
	for {
		time.Sleep(60 * time.Second)

		log.Println("Running scheduled task: ScrapeHeadlines")
		data.ScrapeAndStoreHeadlines(data.Portals)
	}
}
