package scheduler

import (
	"log"
	"time"

	"github.com/farofadev/headlinesbr/model"
)

func Run() {
	for {
		time.Sleep(120 * time.Second)

		log.Println("Running scheduled task: ScrapeHeadlines")
		model.ScrapeAndStoreHeadlines(model.Portals)
	}
}
