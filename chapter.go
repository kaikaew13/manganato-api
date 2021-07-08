package manganatoapi

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Chapter struct {
	ID          string
	MangaID     string
	ChapterName string
	Views       string
	Uploaded    string
	Pages       []Page
}

func createChapterList(m *Manga) {

	c.OnHTML(".row-content-chapter li.a-h", func(h *colly.HTMLElement) {
		ch := Chapter{}

		ch.getChapterID(h.ChildAttr("a.chapter-name", "href"))
		ch.MangaID = m.ID
		ch.ChapterName = h.ChildText("a.chapter-name")
		ch.Views = h.ChildText("span.chapter-view")
		ch.Uploaded = h.ChildText("span.chapter-time")

		m.ChapterList = append(m.ChapterList, ch)

	})
}

func (ch *Chapter) OpenChapterByID() {

	// call createPages() which will do the web scraping
	ch.Pages = createPages(ch.getChapterURL())
}

func (ch *Chapter) getChapterID(url string) {
	ch.ID = getID(url, "-")
}

func (ch *Chapter) getChapterURL() string {
	return fmt.Sprintf("%s%s/chapter-%s", specificMangaURL, ch.MangaID, ch.ID)
}
