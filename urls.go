package manganatoapi

const (
	manganatoURL              string = "manganato.com"
	manganatoURLWithHTTPS            = "https://" + manganatoURL
	readManganatoURL                 = "read" + manganatoURL
	readManganatoURLWihtHTTPS        = "https://" + readManganatoURL
	searchMangaURL                   = manganatoURLWithHTTPS + "/search/story/"
	specificMangaURL                 = readManganatoURLWihtHTTPS + "/manga-"
	searchMangaByAuthorURL           = manganatoURLWithHTTPS + "/author/story/"
	searchMangaByGenreURL            = manganatoURLWithHTTPS + "/genre-"
)
