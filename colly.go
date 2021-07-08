package manganatoapi

import "github.com/gocolly/colly"

var c *colly.Collector

func Setup() {
	c = colly.NewCollector(
		colly.AllowedDomains(
			manganatoURL,
			readManganatoURL,
		),
	)
}
