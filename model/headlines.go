package model

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/farofadev/headlinesbr/config"
	"github.com/farofadev/headlinesbr/database"
	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DefaultPerPage int64 = 25
var Portals = getPortals()

func ScrapeAndStoreHeadlines(portals []Portal) *[]Post {
	results := ScrapeHeadlines(portals)

	StoreHeadlines(*results)

	return results
}

// ScrapeHeadlines Scrape headlines from web
func ScrapeHeadlines(portals []Portal) *[]Post {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	total := len(Portals)
	posts := make([]Post, 0, total)

	wg.Add(total)

	for index := range portals {
		go func(i int) {
			portal := &portals[i]
			collector := colly.NewCollector()
			collector.SetRequestTimeout(30 * time.Second)

			collector.OnHTML(portal.HeadlineSelector, func(element *colly.HTMLElement) {
				link := element.DOM

				if !element.DOM.Is("a") {
					link = element.DOM.Closest("a")
				}

				href := link.AttrOr("href", "")

				if !strings.HasPrefix(href, "http://") && !strings.HasPrefix(href, "https://") {
					href = strings.TrimRight(portal.Url, "/") + "/" + strings.TrimLeft(href, "/")
				}

				post := Post{
					PortalId:  portal.Id,
					Title:     strings.Trim(element.Text, "\n\t\r "),
					Url:       href,
					CreatedAt: time.Now(),
				}

				mutex.Lock()
				posts = append(posts, post)
				mutex.Unlock()
			})

			collector.Visit(portal.Url)
			wg.Done()
		}(index)
	}

	wg.Wait()

	return &posts
}

// StoreHeadlines Save headlines into database
func StoreHeadlines(posts []Post) {
	ctx := context.Background()
	client, _ := database.GetClient(ctx)

	defer client.Disconnect(ctx)

	collection := client.Database(config.Database).Collection("posts")

	for i := range posts {
		result := collection.FindOne(ctx, bson.M{"url": posts[i].Url})

		if result.Err() != nil {
			collection.InsertOne(ctx, posts[i])
		}
	}
}

func PageOffset(page, perPage int64) int64 {
	if page < 1 {
		return 0
	}

	return (page - 1) * perPage
}

func ParsePageAndPerPage(pageStr string, perPageStr string) (int64, int64) {
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	perPage, _ := strconv.ParseInt(perPageStr, 10, 64)

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = DefaultPerPage
	}

	return page, perPage
}

func FetchHeadlines(filters interface{}, findOptions *options.FindOptions) *[]Post {
	ctx := context.Background()
	client, _ := database.GetClient(ctx)

	defer client.Disconnect(ctx)

	collection := client.Database(config.Database).Collection("posts")

	results, _ := collection.Find(
		ctx,
		filters,
		findOptions,
	)

	defer results.Close(ctx)

	posts := []Post{}

	results.All(ctx, &posts)

	for i := range posts {
		posts[i].Portal = FindPortalById(&Portals, posts[i].PortalId)
	}

	return &posts
}

var portalsCache = make(map[uint]*Portal)

func FindPortalById(portals *[]Portal, id uint) *Portal {
	if portal, found := portalsCache[id]; found {
		return portal
	}

	for _, portal := range *portals {
		portalsCache[id] = &portal
		if portal.Id == id {
			return &portal
		}
	}
	return nil
}
