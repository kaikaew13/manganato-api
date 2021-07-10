package manganatoapi

import (
	"errors"

	"github.com/gocolly/colly"
)

var ErrPageNotFound = errors.New("this page does not exist or has been deleted")

var c *colly.Collector

type Searcher struct {
	MethodsDescription map[string]string
}

type Searchable interface {
	getID(string)
}

func NewSearcher() Searcher {
	c = colly.NewCollector(
		colly.AllowedDomains(
			manganatoURL,
			readManganatoURL,
		),
		colly.MaxDepth(2),
	)

	// description of each methods are to be added
	methodDescription := make(map[string]string)

	return Searcher{
		MethodsDescription: methodDescription,
	}
}

func (s *Searcher) SearchManga(name string) (*[]Manga, error) {
	tmp := getMangaList(changeSpaceToUnderscore(name))

	if len(tmp) == 0 {
		return nil, ErrPageNotFound
	}

	return &tmp, nil
}

func (s *Searcher) PickManga(id string) (*Manga, error) {
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

func (s *Searcher) ReadMangaChapter(mangaId, chapterId string) (*[]Page, error) {
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

func (s *Searcher) PickAuthor(authorId string) (*[]Manga, error) {
	a := Author{
		ID: authorId,
	}

	a.getMangaListByAuthorID()

	if len(a.Mangas) == 0 {
		return nil, ErrPageNotFound
	}

	return &a.Mangas, nil
}

func (s *Searcher) PickGenre(genreId string) (*[]Manga, error) {
	g := Genre{
		ID: genreId,
	}

	g.getMangaListByGenreID()

	if len(g.Mangas) == 0 {
		return nil, ErrPageNotFound
	}

	return &g.Mangas, nil
}

func (s *Searcher) SearchLatestUpdatedManga() (*[]Manga, error) {
	tmp := getLatestUpdatedManga()

	if len(tmp) == 0 {
		return nil, ErrPageNotFound
	}

	return &tmp, nil
}

func (s *Searcher) IsSearchable(any interface{}) bool {
	switch any.(type) {
	case Searchable:
		return true
	}
	return false
}
