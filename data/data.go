package data

import (
	"context"
	"fmt"
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

type Portal struct {
	Id               uint   `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description,omitempty" bson:"description,omitempty"`
	Color            string `json:"color,omitempty" bson:"color,omitempty"`
	Url              string `json:"url" bson:"url"`
	HeadlineSelector string `json:"-" bson:"-"`
}

type Post struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty"`
	PortalId  uint      `json:"portal_id" bson:"portal_id"`
	Title     string    `json:"title" bson:"title"`
	Url       string    `json:"url" bson:"url"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Portal    *Portal   `json:"portal,omitempty" bson:"-"`
}

var DefaultPerPage int64 = 25

var Portals = []Portal{
	{
		Id:               1,
		Name:             "Globo",
		Url:              "https://globo.com",
		Color:            "blue",
		HeadlineSelector: "body > section.highlight-container.hui-container.hui-spacing > div.highlight-container__left-area.highlight-main > div > a > h3",
	},
	{
		Id:               2,
		Name:             "UOL",
		Color:            "brown",
		Url:              "https://uol.com.br",
		HeadlineSelector: "#corpo > div:nth-child(1) > div > div.topo-hibrido-central.centraliza.clearfix.bloco-editorial-topo-1 > div.topo-hibrido-hardnews > div.topo-hibrido-hardnews-destaque > div.mod-hibrido-manchete.area-default.manchete-editorial > a > h1",
	},
	{
		Id:               3,
		Name:             "Folha",
		Color:            "black",
		Url:              "https://www.folha.uol.com.br",
		HeadlineSelector: ".c-main-headline__title",
	},
	{
		Id:               4,
		Name:             "Veja",
		Color:            "red",
		Url:              "https://veja.com.br",
		HeadlineSelector: "body > main > section.block.hard-news.light > div > div > div.col-s-12.col-l-9 > div.card.d > a:nth-child(1) > h2",
	},
	{
		Id:               5,
		Name:             "Estadão",
		Color:            "blue",
		Url:              "https://www.estadao.com.br/",
		HeadlineSelector: "#wrapper > section.breaking-news > div > div > div:nth-child(1) > article > div > div > div > div.intro > a > h3",
	},
	{
		Id:               6,
		Name:             "Terra",
		Color:            "orange",
		Url:              "https://www.terra.com.br/",
		HeadlineSelector: "div.card-premium__left > h2 > a",
	},
	{
		Id:               7,
		Name:             "Metro1",
		Color:            "yellow",
		Url:              "https://www.metro1.com.br/",
		HeadlineSelector: "body > div > main section:nth-child(1) article div a",
	},
	{
		Id:               8,
		Name:             "Exame",
		Color:            "red",
		Url:              "https://exame.com",
		HeadlineSelector: "#abril_home_box_widget-101 > div.widget-box.widget-home-box.widget-box- > div.hide_thumb.widget-home-box-list-item.type-post.item-size-g > div > a.widget-home-box-list-item-title > h2",
	},
	{
		Id:               9,
		Name:             "Correio",
		Color:            "red",
		Url:              "https://correio24horas.com.br",
		HeadlineSelector: "#CW9220 > div.destaque-responsivo__container-info > div > a",
	},
	{
		Id:               10,
		Name:             "Metrópoles",
		Color:            "yellow",
		Url:              "https://www.metropoles.com",
		HeadlineSelector: "#m-main > section.m-top-news > div > div > div.column.is-three-quarters-widescreen.is-full > article > div:nth-child(2) > div > h2 > a",
	},
	{
		Id:               11,
		Name:             "BNews",
		Color:            "red",
		Url:              "https://www.bnews.com.br/",
		HeadlineSelector: "h3.title a",
	},
}

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

	wg := sync.WaitGroup{}

	wg.Add(len(posts))
	go func() {
		for i := range posts {
			posts[i].Portal = FindPortalById(&Portals, posts[i].PortalId)
			wg.Done()
		}
	}()
	wg.Wait()

	return &posts
}

var portalsCache = make(map[uint]*Portal)
var portalsCacheLoaded = false

func FindPortalById(portals *[]Portal, id uint) *Portal {
	if portal, found := portalsCache[id]; found {
		return portal
	}

	if portalsCacheLoaded {
		return &Portal{}
	}

	for _, portal := range *portals {
		portalsCache[id] = &portal

		if portal.Id == id {
			return &portal
		}
	}

	portalsCacheLoaded = true

	return &Portal{}
}
