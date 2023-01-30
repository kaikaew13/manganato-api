package manganatoapi

import (
	"fmt"
	"testing"
)

const id string = "dn980422"

// testing main apis

func TestSearchManga(t *testing.T) {
	fmt.Println("on TestSearchManga")
	for i := 0; i < 5; i++ {
		s := NewSearcher()

		mangaName := "chainsaw man"

		mgs, err := s.SearchManga(mangaName)
		if err != nil {
			t.Fatal("not expect to have error")
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
}

// they update the website and manga info is not the same
// func TestPickManga(t *testing.T) {
// 	fmt.Println("on TestPickManga")
// 	for i := 0; i < 5; i++ {
// 		s := NewSearcher()

// 		_, err := s.PickManga(id)
// 		if err != nil {
// 			t.Fatal("not expect to have error")
// 		}

// 		m, err := s.PickManga(id)
// 		if err != nil {
// 			t.Fatal("not expect to have error 2")
// 		}

// 		want := struct {
// 			Name         string
// 			Author       string
// 			Alternatives string
// 			Length       int
// 			MangaID      string
// 		}{
// 			Name:         "Chainsaw Man",
// 			Author:       "Tatsuki Fujimoto",
// 			Alternatives: "Chainsawman, チェンソーマン",
// 			Length:       97,
// 			MangaID:      id,
// 		}

// 		compareNameHelper(t, m.Name, want.Name)
// 		compareAuthorHelper(t, m.Author.Name, want.Author)
// 		compareAlternativesHelper(t, m.Alternatives, want.Alternatives)
// 		compareChaptersHelper(t, m.Chapters, struct {
// 			Length  int
// 			MangaID string
// 		}{
// 			Length:  want.Length,
// 			MangaID: want.MangaID,
// 		})
// 	}
// }

func TestReadMangaChapter(t *testing.T) {
	fmt.Println("on TestReadMangaChapter")
	for i := 0; i < 5; i++ {
		s := NewSearcher()

		pgs, err := s.ReadMangaChapter(id, "97")
		if err != nil {
			t.Fatal("not expect to have error")
		}

		want := struct {
			Length       int
			FirstPageURL string
		}{
			Length:       23,
			FirstPageURL: "https://v11.mkklcdnv6tempv4.com/img/tab_11/02/90/65/dn980422/chapter_97_love_love_chainsaw/1-n.jpg",
		}
		comparePagesHelper(t, *pgs, want)
	}
}

func TestSearchMangaByAuthor(t *testing.T) {
	fmt.Println("on TestSearchMangaByAuthor")
	for i := 0; i < 5; i++ {
		s := NewSearcher()

		mgs, err := s.PickAuthor("fHx0YXRzdWtpX2Z1amltb3Rv")
		if err != nil {
			t.Fatal("not expect to have error")
		}

		want := struct {
			Length int
			Author string
		}{
			Length: 7,
			Author: "Fujimoto Tatsuki",
		}

		compareMangasHelper(t, len(*mgs), want.Length)
		compareAuthorHelper(t, (*mgs)[1].Author.Name, want.Author)
	}
}

func TestSearchMangaByGenre(t *testing.T) {
	fmt.Println("on TestSearchMangaByGenre")
	for i := 0; i < 5; i++ {
		s := NewSearcher()

		mgs, err := s.PickGenre("2")
		if err != nil {
			t.Fatal("not expect to have error")
		}

		want := struct {
			Length int
		}{
			Length: 24,
		}

		compareMangasHelper(t, len(*mgs), want.Length)
	}
}

func TestSearchTopManga(t *testing.T) {
	fmt.Println("on TestSearchTopManga")
	for i := 0; i < 5; i++ {
		s := NewSearcher()

		mgs, err := s.SearchLatestUpdatedManga()
		if err != nil {
			t.Fatalf("not expect to have error: %s", err.Error())
		}

		want := struct {
			Length int
		}{
			Length: 56,
		}

		compareMangasHelper(t, len(*mgs), want.Length)
	}
}

func TestNotFound(t *testing.T) {
	fmt.Println("on TestNotFound")
	s := NewSearcher()

	_, err := s.SearchManga(" asdlfjas j laja j")
	notFoundHelper(t, err)

	_, err = s.PickManga("to70571")
	notFoundHelper(t, err)

	_, err = s.ReadMangaChapter("to70571", "1")
	notFoundHelper(t, err)

	_, err = s.ReadMangaChapter("to970571", "1000")
	notFoundHelper(t, err)

	// for https://manganato.com/author/story/ route, short random string does not
	// result in 404 error, only long strings or string with more than one consecutive
	// space will result in 404 error
	// case one: with long string
	_, err = s.PickAuthor("asldjfsjflsajfljdsafljasdfljafjaslfjsfldsjflsdjfkjflsjljsfjdaflfjjsdaljs")
	notFoundHelper(t, err)

	// case two: with more than one consecutive space
	_, err = s.PickAuthor("a  b")
	notFoundHelper(t, err)

	_, err = s.PickGenre("abc")
	notFoundHelper(t, err)
}

func TestIsSearchable(t *testing.T) {
	fmt.Println("on TestIsSearchable")
	s := NewSearcher()

	got := s.IsSearchable(1)
	want := false
	compareIsSearchableHelper(t, got, want)

	got = s.IsSearchable(&Manga{})
	want = true
	compareIsSearchableHelper(t, got, want)

	got = s.IsSearchable(&Author{})
	want = true
	compareIsSearchableHelper(t, got, want)
}

//

func TestGetChapterURL(t *testing.T) {
	fmt.Println("on TestGetChapterURL")
	ch := Chapter{
		ID:      "97",
		MangaID: id,
	}

	url := ch.getChapterURL()

	want := specificMangaURL + id + "/chapter-97"

	compareURLHelper(t, url, want)
}

func TestCreatePages(t *testing.T) {
	fmt.Println("on TestCreatePages")
	initCrawler()
	defer deleteCrawler()

	pgs := createPages("https://chapmanganato.com/manga-dn980422/chapter-97")
	if len(pgs) != 23 {
		t.Error("suppose to have 23 pages, got ", len(pgs))
	}

	want := struct {
		Length       int
		FirstPageURL string
	}{
		Length:       23,
		FirstPageURL: "https://v11.mkklcdnv6tempv4.com/img/tab_11/02/90/65/dn980422/chapter_97_love_love_chainsaw/1-n.jpg",
	}

	comparePagesHelper(t, pgs, want)
}

func notFoundHelper(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("should have error of type ErrPageNotFound")
	}
	if err != ErrPageNotFound {
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

func compareIsSearchableHelper(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("wanted %t, got %t", want, got)
	}
}
