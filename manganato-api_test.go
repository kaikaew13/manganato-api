package manganatoapi

import (
	"testing"
)

const id string = "dn980422"

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

	m := Manga{
		ID: id,
	}

	m.SearchMangaByID()

	want := struct {
		Name         string
		Author       string
		Alternatives string
	}{
		Name:         "Chainsaw Man",
		Author:       "Tatsuki Fujimoto",
		Alternatives: "Chainsawman, チェンソーマン",
	}

	if m.Name != want.Name {
		t.Errorf("wanted manga with name %s, got %s", want.Name, m.Name)
	}
	if m.Author != want.Author {
		t.Errorf("wanted manga with author %s, got %s", want.Author, m.Author)
	}
	if m.Alternatives != want.Alternatives {
		t.Errorf("wanted manga with alternatives %s, got %s", want.Alternatives, m.Alternatives)
	}
}

func TestCreateChapterList(t *testing.T) {
	Setup()

	chapters := createChapterList(id)

	want := 97

	if len(chapters) != want {
		t.Errorf("wanted manga to have %d chapters, got %d", want, len(chapters))
	}
}

func TestCreateAuthor(t *testing.T) {
	Setup()

	createAuthor(id)
}
