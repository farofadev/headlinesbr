package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/farofadev/headlinesbr/data"
	"github.com/farofadev/headlinesbr/database"
	"github.com/gocolly/colly"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/headlines", FetchHeadlines)

	http.ListenAndServe(":8080", router)

}

func FetchHeadlines(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	results := ScrapeHeadlines(data.Portals)

	StoreHeadlines(*results)

	http.Redirect(w, r, "/", http.StatusSeeOther)

	// w.Header().Set("Location", "/")

	// w.Write([]byte(""))

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx := context.Background()
	client := database.GetClient()

	postsCollection := client.Database("headlinesbr").Collection("posts")

	results, _ := postsCollection.Find(ctx, bson.M{})

	bsonM := []bson.M{}

	results.All(ctx, &bsonM)

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(bsonM)

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
	client := database.GetClient()

	postsCollection := client.Database("headlinesbr").Collection("posts")

	for i := range posts {

		result := postsCollection.FindOne(ctx, bson.M{"url": posts[i].Url})
		var p data.Post
		err := result.Decode(&p)
		if err != nil {
			postsCollection.InsertOne(ctx, posts[i].BsonM())
		}

	}

}
