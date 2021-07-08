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

	SearchMangaByID("dn980422")
}
