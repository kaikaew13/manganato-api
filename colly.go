package manganatoapi

import (
	"github.com/gocolly/colly"
)

var c *colly.Collector

func InitCrawler() {
	c = colly.NewCollector(
		colly.AllowedDomains(
			manganatoURL,
			readManganatoURL,
		),
		colly.MaxDepth(2),
	)
}
