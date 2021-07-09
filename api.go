package manganatoapi

func SearchManga(name string) (*[]Manga, error) {
	tmp := getMangaList(changeSpaceToUnderscore(name))

	if len(tmp) == 0 {
		return nil, newNotFoundError()
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
		return nil, newNotFoundError()
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
		return nil, newNotFoundError()
	}

	return &ch.Pages, nil
}

func SearchMangaByAuthor(authorId string) (*[]Manga, error) {
	a := Author{
		ID: authorId,
	}

	a.getMangaListByAuthorID()

	if len(a.Mangas) == 0 {
		return nil, newNotFoundError()
	}

	return &a.Mangas, nil
}

func SearchMangaByGenre(genreId string) (*[]Manga, error) {
	g := Genre{
		ID: genreId,
	}

	g.getMangaListByGenreID()

	if len(g.Mangas) == 0 {
		return nil, newNotFoundError()
	}

	return &g.Mangas, nil
}
