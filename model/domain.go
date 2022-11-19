package model

import "time"

type Locale string

const (
	Brasil Locale = "Brasil"
	Bahia  Locale = "Bahia"
)

type Portal struct {
	Id               uint   `json:"id,omitempty" bson:"_id,omitempty"`
	Locale           Locale `json:"locale"`
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

func getPortals() []Portal {

	return []Portal{
		{
			Id:               1,
			Locale:           Brasil,
			Name:             "globo.com",
			Url:              "https://globo.com",
			Color:            "#2f74aa",
			HeadlineSelector: "body > section.highlight-container.hui-container > div > div > div.headline.post-main.theme-jornalismo.simple.has-bullets > div.headline__container > div > a > h2",
		},
		//{
		//	Id:               2,
		//	Locale:           Brasil,
		//	Name:             "UOL",
		//	Color:            "#444444",
		//	Url:              "https://uol.com.br",
		//	HeadlineSelector: "",
		//},
		{
			Id:               3,
			Locale:           Brasil,
			Name:             "Folha",
			Color:            "#767e84",
			Url:              "https://www.folha.uol.com.br",
			HeadlineSelector: ".c-main-headline__title",
		},
		{
			Id:               4,
			Locale:           Brasil,
			Name:             "Veja",
			Color:            "#d74747",
			Url:              "https://veja.com.br",
			HeadlineSelector: "body > main > section.block.hard-news.light > div > div > div.col-s-12.col-l-9 > div.card.d > a:nth-child(1) > h2",
		},
		{
			Id:               5,
			Locale:           Brasil,
			Name:             "Estadão",
			Color:            "#516c8b",
			Url:              "https://www.estadao.com.br/",
			HeadlineSelector: "#fusion-app > div > div > div:nth-child(1) > div > div.col-12.col-xl-8 > div > div > a > h2",
		},
		{
			Id:               6,
			Locale:           Brasil,
			Name:             "Terra",
			Color:            "#f9772f",
			Url:              "https://www.terra.com.br/",
			HeadlineSelector: "div.card-news__text > h2",
		},
		{
			Id:               7,
			Locale:           Bahia,
			Name:             "Metro1",
			Color:            "#f9772f",
			Url:              "https://www.metro1.com.br/",
			HeadlineSelector: "body > div > main section:nth-child(1) article div a",
		},
		{
			Id:               8,
			Locale:           Brasil,
			Name:             "Exame",
			Color:            "#e52a12",
			Url:              "https://exame.com",
			HeadlineSelector: "#abril_home_box_widget-101 > div.widget-box.widget-home-box.widget-box- > div.hide_thumb.widget-home-box-list-item.type-post.item-size-g > div > a.widget-home-box-list-item-title > h2",
		},
		{
			Id:               9,
			Locale:           Bahia,
			Name:             "Correio",
			Color:            "#002142",
			Url:              "https://correio24horas.com.br",
			HeadlineSelector: "#CW9220 > div.destaque-responsivo__container-info > div > a",
		},
		{
			Id:               10,
			Locale:           Brasil,
			Name:             "Metrópoles",
			Color:            "#c32417",
			Url:              "https://www.metropoles.com",
			HeadlineSelector: "#m-main > section.m-top-news > div > div > div.column.is-three-quarters-widescreen.is-full > article > div:nth-child(2) > div > h2 > a",
		},
		{
			Id:               11,
			Locale:           Bahia,
			Name:             "BNews",
			Color:            "red",
			Url:              "https://www.bnews.com.br/",
			HeadlineSelector: "h3.title a",
		},

		{
			Id:               12,
			Locale:           Brasil,
			Name:             "BBC News Brasil",
			Color:            "red",
			Url:              "https://www.bbc.com/portuguese",
			HeadlineSelector: "#main-wrapper > div > main > div > section:nth-child(1) > div.bbc-1dblbh1.efnv93c1 > ul > li.ebmt73l0.bbc-mhiwdf.e13i2e3d1 > div > div.bbc-14gzkm2.e718b9o0 > h3 > a",
		},
		{
			Id:               13,
			Locale:           Brasil,
			Name:             "R7",
			Color:            "#218EE1",
			Url:              "https://www.r7.com/",
			HeadlineSelector: "#box_61d48b22cd77c0118500051a > div > div > div > div > article > h3 > a",
		},
		{
			Id:               14,
			Locale:           Brasil,
			Name:             "CNN Brasil",
			Color:            "#cc0000",
			Url:              "https://www.cnnbrasil.com.br/",
			HeadlineSelector: "#block1847327 > div > div > a > h2",
		},
	}
}
