package manganatoapi

import (
	"testing"
)

const id string = "dn980422"

// testing main apis

func TestSearchManga(t *testing.T) {
	InitCrawler()

	mangaName := "chainsaw man"

	mgs, err := SearchManga(mangaName)
	if err != nil {
		t.Error("not expect to have error")
	}

	want := struct {
		Length int
		Name   string
		Author string
	}{
		Length: 1,
		Name:   "Chainsaw Man",
		Author: "Tatsuki Fujimoto",
	}

	compareMangasHelper(t, len(*mgs), want.Length)
	compareNameHelper(t, (*mgs)[0].Name, want.Name)
	compareAuthorHelper(t, (*mgs)[0].Author.Name, want.Author)
}

func TestPickManga(t *testing.T) {
	InitCrawler()

	m, err := PickManga(id)
	if err != nil {
		t.Error("not expect to have error")
	}

	want := struct {
		Name         string
		Author       string
		Alternatives string
		Length       int
		MangaID      string
	}{
		Name:         "Chainsaw Man",
		Author:       "Tatsuki Fujimoto",
		Alternatives: "Chainsawman, チェンソーマン",
		Length:       97,
		MangaID:      id,
	}

	compareNameHelper(t, m.Name, want.Name)
	compareAuthorHelper(t, m.Author.Name, want.Author)
	compareAlternativesHelper(t, m.Alternatives, want.Alternatives)
	compareChaptersHelper(t, m.ChapterList, struct {
		Length  int
		MangaID string
	}{
		Length:  want.Length,
		MangaID: want.MangaID,
	})
}

func TestReadMangaChapter(t *testing.T) {
	InitCrawler()

	pgs, err := ReadMangaChapter(id, "97")
	if err != nil {
		t.Error("not expect to have error")
	}

	want := struct {
		Length       int
		FirstPageURL string
	}{
		Length:       23,
		FirstPageURL: "https://s51.mkklcdnv6tempv2.com/mangakakalot/i2/ix917953/chapter_97_love_love_chainsaw/1.jpg",
	}

	comparePagesHelper(t, *pgs, want)
}

func TestSearchMangaByAuthor(t *testing.T) {
	InitCrawler()

	mgs, err := SearchMangaByAuthor("fHx0YXRzdWtpX2Z1amltb3Rv")
	if err != nil {
		t.Error("not expect to have error")
	}

	want := struct {
		Length int
		Name   string
		Author string
	}{
		Length: 7,
		Name:   "Fire Punch",
		Author: "Fujimoto Tatsuki",
	}

	compareMangasHelper(t, len(*mgs), want.Length)
	compareNameHelper(t, (*mgs)[1].Name, want.Name)
	compareAuthorHelper(t, (*mgs)[1].Author.Name, want.Author)
}

func TestSearchMangaByGenre(t *testing.T) {
	InitCrawler()

	mgs, err := SearchMangaByGenre("2")
	if err != nil {
		t.Error("not expect to have error")
	}

	want := struct {
		Length int
	}{
		Length: 24,
	}

	compareMangasHelper(t, len(*mgs), want.Length)
}

func TestSearchTopManga(t *testing.T) {
	InitCrawler()

	mgs, err := SearchLatestUpdatedManga()
	if err != nil {
		t.Errorf("not expect to have error: %s", err.Error())
	}

	want := struct {
		Length int
	}{
		Length: 56,
	}

	compareMangasHelper(t, len(*mgs), want.Length)
}

//

func TestGetChapterURL(t *testing.T) {
	ch := Chapter{
		ID:      "97",
		MangaID: id,
	}

	url := ch.getChapterURL()

	want := specificMangaURL + id + "/chapter-97"

	compareURLHelper(t, url, want)
}

func TestCreatePages(t *testing.T) {
	InitCrawler()

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

func TestNotFound(t *testing.T) {
	InitCrawler()

	_, err := SearchManga(" asdlfjas j laja j")
	notFoundHelper(t, err)

	_, err = PickManga("to70571")
	notFoundHelper(t, err)

	_, err = ReadMangaChapter("to70571", "1")
	notFoundHelper(t, err)

	_, err = ReadMangaChapter("to970571", "1000")
	notFoundHelper(t, err)

	// for https://manganato.com/author/story/ route, short random string does not
	// result in 404 error, only long strings or string with more than one consecutive
	// space will result in 404 error
	// case one: with long string
	_, err = SearchMangaByAuthor("asldjfsjflsajfljdsafljasdfljafjaslfjsfldsjflsdjfkjflsjljsfjdaflfjjsdaljs")
	notFoundHelper(t, err)

	// case two: with more than one consecutive space
	_, err = SearchMangaByAuthor("a  b")
	notFoundHelper(t, err)

	_, err = SearchMangaByGenre("abc")
	notFoundHelper(t, err)

}

func notFoundHelper(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("should have error of type NotFoundError")
	}
	if _, ok := err.(*NotFoundError); !ok {
		t.Error(err.Error())
	}
}

func compareNameHelper(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("wanted manga with name %s, got %s", want, got)
	}
}

func compareAuthorHelper(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("wanted manga with author %s, got %s", want, got)
	}
}

func compareMangasHelper(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("wanted manga list of length %d, got %d", want, got)
	}
}

func compareAlternativesHelper(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("wanted manga with alternatives %s, got %s", want, got)
	}
}

func compareURLHelper(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("wanted url string of %s, got %s", want, got)
	}
}

func compareChaptersHelper(t testing.TB, got []Chapter, want struct {
	Length  int
	MangaID string
}) {
	t.Helper()

	if len(got) != want.Length {
		t.Errorf("wanted manga to have %d chapters, got %d", want.Length, len(got))
	}
	if got[0].MangaID != want.MangaID {
		t.Errorf("wanted chapter to be related with manga with id %s, got %s", want.MangaID, got[0].MangaID)
	}
}

func comparePagesHelper(t testing.TB, got []Page, want struct {
	Length       int
	FirstPageURL string
}) {
	t.Helper()

	if len(got) != want.Length {
		t.Errorf("wanted a chapter with %d pages, got %d", want.Length, len(got))
	}
	compareURLHelper(t, got[0].ImageURL, want.FirstPageURL)
}
