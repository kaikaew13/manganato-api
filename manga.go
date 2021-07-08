package manganatoapi

type Manga struct {
	ID           string
	Name         string
	Alternatives []string
	Authors      []Author
	Status       Status
	Updates      string
	Views        int
	Rating       Rating
	Description  string
	Genres       []Genre
	ChapterList  []Chapter
}
