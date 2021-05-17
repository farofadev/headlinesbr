package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/farofadev/headlinesbr/data"
	"github.com/farofadev/headlinesbr/database"
	"github.com/gocolly/colly"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

var dbname = os.Getenv("MONGO_DATABASE")

func scheduler() {
	for {
		time.Sleep(1 * time.Minute)

		log.Println("Running scheduled task: ScrapeHeadlines")
		go ScrapeHeadlines(data.Portals)
	}
}

func main() {
	go scheduler()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/headlines", FetchHeadlines)

	http.ListenAndServe(":8080", router)
}

func FetchHeadlines(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results := ScrapeHeadlines(data.Portals)

	StoreHeadlines(*results)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := context.Background()
	client, _ := database.GetClient(ctx)

	defer client.Disconnect(ctx)

	collection := client.Database(dbname).Collection("posts")
	results, _ := collection.Find(ctx, bson.M{})

	defer results.Close(ctx)

	posts := []data.Post{}

	results.All(ctx, &posts)

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(posts)
}

// ScrapeHeadlines Scrape headlines from web
func ScrapeHeadlines(portals []data.Portal) *[]data.Post {
	wg := sync.WaitGroup{}
	total := len(data.Portals)
	posts := make([]data.Post, 0, total)

	wg.Add(total)

	for index := range portals {
		go func(i int) {
			portal := &portals[i]
			collector := colly.NewCollector()
			collector.SetRequestTimeout(30 * time.Second)

			collector.OnHTML(portal.HeadlineSelector, func(element *colly.HTMLElement) {
				post := data.Post{
					PortalId:  portal.Id,
					Title:     strings.Trim(element.Text, "\n\t\r "),
					Url:       element.DOM.Closest("a").AttrOr("href", ""),
					CreatedAt: time.Now(),
				}

				posts = append(posts, post)
			})

			collector.Visit(portal.Url)
			wg.Done()
		}(index)
	}

	wg.Wait()

	return &posts
}

// StoreHeadlines Save headlines into database
func StoreHeadlines(posts []data.Post) {
	ctx := context.Background()
	client, _ := database.GetClient(ctx)

	defer client.Disconnect(ctx)

	collection := client.Database(dbname).Collection("posts")

	for i := range posts {
		result := collection.FindOne(ctx, bson.M{"url": posts[i].Url})

		if result.Err() != nil {
			collection.InsertOne(ctx, posts[i])
		}
	}
}
