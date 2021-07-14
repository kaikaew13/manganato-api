package manganatoapi

import (
	"errors"
	"sync"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// throws when no results are found
var ErrPageNotFound = errors.New("this page does not exist or has been deleted")

var c *colly.Collector

// provides methods for fetching data from https://manganato.com
type Searcher struct {
	MethodsDescription map[string]string
}

//  any struct types with ID implement Searchable
type Searchable interface {
	getID(string)
}

func initCrawler() {
	c = colly.NewCollector(
		colly.AllowedDomains(
			manganatoURL,
			readManganatoURL,
		),
		colly.MaxDepth(2),
		colly.Async(true),
		// colly.Debugger(&debug.LogDebugger{}),
	)

	extensions.RandomUserAgent(c)
}

func cloneCrawler() *colly.Collector {
	c2 := c.Clone()
	extensions.RandomUserAgent(c2)
	return c2
}

func deleteCrawler() {
	c = nil
}

// returns a new Searcher
func NewSearcher() Searcher {
	methodDescription := map[string]string{
		"SearchManga":              "receives name of a manga user wants to search for and returns a list of mangas that match the name",
		"PickManga":                "receives the id of the specific manga then returns that manga if found",
		"ReadMangaChapter":         "receives the manga id and chapter id then returns pages of that specific chapter",
		"PickAuthor":               "receives the id of the author then returns a list of mangas by him/her",
		"PickGenre":                "receives genre id then returns a list of mangas with that genre",
		"SearchLatestUpdatedManga": "returns list of latest updated mangas from the first page of https://manganato.com",
		"IsSearchable":             "returns whether the struct type implements Searchable interface",
	}

	return Searcher{
		MethodsDescription: methodDescription,
	}
}

// receives name of a manga user wants to search for
// and returns a list of mangas that match the name
//
// https://manganato.com/search/story/<name>
func (s *Searcher) SearchManga(name string) (*[]Manga, error) {
	initCrawler()
	defer deleteCrawler()

	tmp := getMangaList(changeSpaceToUnderscore(name))

	if len(tmp) == 0 {
		return nil, ErrPageNotFound
	}

	var wg sync.WaitGroup

	mgs := []Manga{}

	for _, mg := range tmp {
		wg.Add(1)

		go func(m Manga, c2 *colly.Collector) {
			defer wg.Done()

			createAuthor(&m, c2)
			mgs = append(mgs, m)

		}(mg, cloneCrawler())
	}

	wg.Wait()

	return &mgs, nil
}

// receives the id of the specific manga
// then returns that manga if found
//
// https://readmanganato.com/manga-<id>
func (s *Searcher) PickManga(id string) (*Manga, error) {
	initCrawler()
	defer deleteCrawler()

	m := Manga{
		ID: id,
	}
	tmp := m

	m.getMangaByID()

	if m.compareManga(&tmp) {
		return nil, ErrPageNotFound
	}

	return &m, nil
}

// receives the manga id and chapter id
// then returns pages of that specific chapter
//
// https://readmanganato.com/manga-<mangaId>/chapter-<chapterId>
func (s *Searcher) ReadMangaChapter(mangaId, chapterId string) (*[]Page, error) {
	initCrawler()
	defer deleteCrawler()

	ch := Chapter{
		ID:      chapterId,
		MangaID: mangaId,
	}

	ch.getChapterByID()

	if len(ch.Pages) == 0 {
		return nil, ErrPageNotFound
	}

	return &ch.Pages, nil
}

// receives the id of the author then returns a list of mangas by him/her
//
// https://manganato.com/author/story/<authorId>
func (s *Searcher) PickAuthor(authorId string) (*[]Manga, error) {
	initCrawler()
	defer deleteCrawler()

	a := Author{
		ID: authorId,
	}
	a.getMangaListByAuthorID()

	if len(a.Mangas) == 0 {
		return nil, ErrPageNotFound
	}

	var wg sync.WaitGroup

	for i, mg := range a.Mangas {
		wg.Add(1)
		go func(m Manga, index int, c2 *colly.Collector) {
			defer wg.Done()

			createAuthor(&m, c2)
			a.Mangas[index] = m
		}(mg, i, cloneCrawler())
	}

	wg.Wait()

	return &a.Mangas, nil
}

// receives genre id then returns a list of mangas with that genre
//
// https://manganato.com/genre-<genreId>
func (s *Searcher) PickGenre(genreId string) (*[]Manga, error) {
	initCrawler()
	defer deleteCrawler()

	g := Genre{
		ID: genreId,
	}

	g.getMangaListByGenreID()

	if len(g.Mangas) == 0 {
		return nil, ErrPageNotFound
	}

	return &g.Mangas, nil
}

// returns list of latest updated mangas
// from the first page of https://manganato.com
func (s *Searcher) SearchLatestUpdatedManga() (*[]Manga, error) {
	initCrawler()
	defer deleteCrawler()

	tmp := getLatestUpdatedManga()

	if len(tmp) == 0 {
		return nil, ErrPageNotFound
	}

	return &tmp, nil
}

// returns whether the struct type implements Searchable interface
func (s *Searcher) IsSearchable(any interface{}) bool {
	switch any.(type) {
	case Searchable:
		return true
	}
	return false
}
