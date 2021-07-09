package manganatoapi

import (
	"testing"
)

const id string = "dn980422"

func TestSearchManga(t *testing.T) {
	Setup()

	mangaName := "chainsaw man"

	mangas := SearchManga(mangaName)

	want := struct {
		Length int
		Name   string
		Author string
	}{
		Length: 1,
		Name:   "Chainsaw Man",
		Author: "Tatsuki Fujimoto",
	}

	// to be refactored
	if len(mangas) != want.Length {
		t.Errorf("wanted slice of length %d, got %d", want.Length, len(mangas))
	}
	if mangas[0].Name != want.Name {
		t.Errorf("wanted manga with name %s, got %s", want.Name, mangas[0].Name)
	}
	if mangas[0].Author.Name != want.Author {
		t.Errorf("wanted manga with author %s, got %s", want.Author, mangas[0].Author.Name)
	}
}

func TestPickManga(t *testing.T) {
	Setup()

	m := PickManga(id)

	want := struct {
		Name         string
		Author       string
		Alternatives string
		Chapters     int
		MangaID      string
	}{
		Name:         "Chainsaw Man",
		Author:       "Tatsuki Fujimoto",
		Alternatives: "Chainsawman, チェンソーマン",
		Chapters:     97,
		MangaID:      id,
	}

	if m.Name != want.Name {
		t.Errorf("wanted manga with name %s, got %s", want.Name, m.Name)
	}
	if m.Author.Name != want.Author {
		t.Errorf("wanted manga with author %s, got %s", want.Author, m.Author.Name)
	}
	if m.Alternatives != want.Alternatives {
		t.Errorf("wanted manga with alternatives %s, got %s", want.Alternatives, m.Alternatives)
	}
	if len(m.ChapterList) != want.Chapters {
		t.Errorf("wanted manga to have %d chapters, got %d", want.Chapters, len(m.ChapterList))
	}
	if m.ChapterList[0].MangaID != want.MangaID {
		t.Errorf("wanted chapter to be related with manga with id %s, got %s", want.MangaID, m.ChapterList[0].MangaID)
	}
}

func TestGetChapterURL(t *testing.T) {
	ch := Chapter{
		ID:      "97",
		MangaID: id,
	}

	url := ch.getChapterURL()

	want := specificMangaURL + id + "/chapter-97"

	if url != want {
		t.Errorf("wanted url string of %s, got %s", want, url)
	}
}

func TestCreatePages(t *testing.T) {
	Setup()

	pgs := createPages("https://readmanganato.com/manga-dn980422/chapter-97")

	want := struct {
		Length       int
		FirstPageURL string
	}{
		Length:       23,
		FirstPageURL: "https://s51.mkklcdnv6tempv2.com/mangakakalot/i2/ix917953/chapter_97_love_love_chainsaw/1.jpg",
	}

	comparePagesHelper(t, pgs, want)
}

func TestReadMangaChapter(t *testing.T) {
	Setup()

	pgs := ReadMangaChapter(id, "97")

	want := struct {
		Length       int
		FirstPageURL string
	}{
		Length:       23,
		FirstPageURL: "https://s51.mkklcdnv6tempv2.com/mangakakalot/i2/ix917953/chapter_97_love_love_chainsaw/1.jpg",
	}

	comparePagesHelper(t, pgs, want)
}

func comparePagesHelper(t testing.TB, got []Page, want struct {
	Length       int
	FirstPageURL string
}) {
	t.Helper()

	if len(got) != want.Length {
		t.Errorf("wanted a chapter with %d pages, got %d", want.Length, len(got))
	}
	if got[0].ImageURL != want.FirstPageURL {
		t.Errorf("wanted url of the first page to be %s, got %s", want.FirstPageURL, got[0].ImageURL)
	}
}
