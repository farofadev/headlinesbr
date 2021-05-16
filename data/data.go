package data

type Portal struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Url              string `json:"url"`
	HeadlineSelector string `json:"-"`
	Headline         string `json:"headline"`
}

var Portals []Portal

func init() {
	Portals = []Portal{
		Globo,
		Uol,
		Folha,
		Veja,
		Estadao,
	}
}

var Globo = Portal{
	Id:               1,
	Name:             "Globo.com",
	Url:              "https://globo.com",
	HeadlineSelector: "body > section.highlight-container.hui-container.hui-spacing > div.highlight-container__left-area.highlight-main.headline > div > a > h3",
}

var Uol = Portal{
	Id:               2,
	Name:             "UOL",
	Url:              "https://uol.com.br",
	HeadlineSelector: "#corpo > div:nth-child(1) > div > div.topo-hibrido-central.centraliza.clearfix.bloco-editorial-topo-1 > div.topo-hibrido-hardnews > div.topo-hibrido-hardnews-destaque > div.mod-hibrido-manchete.area-default.manchete-editorial > a > h1",
}

var Folha = Portal{
	Id:               3,
	Name:             "Folha de São Paulo",
	Url:              "https://www.folha.uol.com.br",
	HeadlineSelector: ".c-main-headline__title",
}

var Veja = Portal{
	Id:               4,
	Name:             "Revista Veja",
	Url:              "https://veja.com.br",
	HeadlineSelector: "body > main > section.block.hard-news.light > div > div > div.col-s-12.col-l-9 > div.card.d > a:nth-child(1) > h2",
}

var Estadao = Portal{
	Id:               5,
	Name:             "Jornal O Estado de São Paulo",
	Url:              "https://www.estadao.com.br/",
	HeadlineSelector: "#wrapper > section.breaking-news > div > div > div:nth-child(1) > article > div > div > div > div.intro > a > h3",
}
