package manganatoapi

func SearchManga(name string) []Manga {
	return getMangaList(changeSpaceToUnderscore(name))
}

func PickManga(id string) Manga {
	m := Manga{
		ID: id,
	}

	m.getMangaByID()

	return m
}

func ReadMangaChapter(mangaId, chapterId string) []Page {
	ch := Chapter{
		ID:      chapterId,
		MangaID: mangaId,
	}

	ch.getChapterByID()

	return ch.Pages
}
