package manganatoapi

import "errors"

var ErrPageNotFound = errors.New("this page does not exist or has been deleted")

func SearchManga(name string) (*[]Manga, error) {
	tmp := getMangaList(changeSpaceToUnderscore(name))

	if len(tmp) == 0 {
		return nil, ErrPageNotFound
	}

	return &tmp, nil
}

func PickManga(id string) (*Manga, error) {
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

func ReadMangaChapter(mangaId, chapterId string) (*[]Page, error) {
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

func SearchMangaByAuthor(authorId string) (*[]Manga, error) {
	a := Author{
		ID: authorId,
	}

	a.getMangaListByAuthorID()

	if len(a.Mangas) == 0 {
		return nil, ErrPageNotFound
	}

	return &a.Mangas, nil
}

func SearchMangaByGenre(genreId string) (*[]Manga, error) {
	g := Genre{
		ID: genreId,
	}

	g.getMangaListByGenreID()

	if len(g.Mangas) == 0 {
		return nil, ErrPageNotFound
	}

	return &g.Mangas, nil
}

func SearchLatestUpdatedManga() (*[]Manga, error) {
	tmp := getLatestUpdatedManga()

	if len(tmp) == 0 {
		return nil, ErrPageNotFound
	}

	return &tmp, nil
}
