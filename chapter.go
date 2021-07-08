package manganatoapi

import (
	"github.com/gocolly/colly"
)

type Chapter struct {
	ID          string
	ChapterName string
	Views       string
	Uploaded    string
	// Pages
}

func createChapterList(id string) []Chapter {
	chapters := []Chapter{}

	c.OnHTML(".row-content-chapter li.a-h", func(h *colly.HTMLElement) {

		id := getId(h.ChildAttr("a.chapter-name", "href"))
		chapterName := h.ChildText("a.chapter-name")
		views := h.ChildText("span.chapter-view")
		uploaded := h.ChildText("span.chapter-time")

		chapters = append(chapters, Chapter{
			ID:          id,
			ChapterName: chapterName,
			Views:       views,
			Uploaded:    uploaded,
		})

	})

	c.Visit(specificMangaURL + id)

	return chapters
}
