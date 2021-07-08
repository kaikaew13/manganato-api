package manganatoapi

import (
	"testing"
)

func TestSearchManga(t *testing.T) {
	Setup()

	mangaName := "chainsaw man"
	mangaNameFormatted := changeSpaceToUnderscore(mangaName)

	mangas := SearchManga(mangaNameFormatted)

	want := struct {
		Length int
		Name   string
		Author string
	}{
		Length: 1,
		Name:   "Chainsaw Man",
		Author: "Tatsuki Fujimoto",
	}

	if len(mangas) != want.Length {
		t.Errorf("wanted slice of length %d, got %d", want.Length, len(mangas))
	}
	if mangas[0].Name != want.Name {
		t.Errorf("wanted manga with name %s, got %s", want.Name, mangas[0].Name)
	}
	if mangas[0].Author != want.Author {
		t.Errorf("wanted manga with author %s, got %s", want.Author, mangas[0].Author)
	}
}

func TestSearchMangaByID(t *testing.T) {
	Setup()

	manga := SearchMangaByID("dn980422")

	want := struct {
		Name         string
		Author       string
		Alternatives string
	}{
		Name:         "Chainsaw Man",
		Author:       "Tatsuki Fujimoto",
		Alternatives: "Chainsawman, チェンソーマン",
	}

	if manga.Name != want.Name {
		t.Errorf("wanted manga with name %s, got %s", want.Name, manga.Name)
	}
	if manga.Author != want.Author {
		t.Errorf("wanted manga with author %s, got %s", want.Author, manga.Author)
	}
	if manga.Alternatives != want.Alternatives {
		t.Errorf("wanted manga with alternatives %s, got %s", want.Alternatives, manga.Alternatives)
	}
}

func TestCreateChapterList(t *testing.T) {
	Setup()

	chapters := createChapterList("dn980422")

	want := 97

	if len(chapters) != want {
		t.Errorf("wanted manga to have %d chapters, got %d", want, len(chapters))
	}
}
