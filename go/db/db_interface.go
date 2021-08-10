package db

import (
	model "github.com/lea-video/backend/go/model"
)

type DBInterface interface {
	FetchIDs() []string

	CreateAlbum(*model.Album)
	FetchAlbum(string) *model.RawAlbum

	CreateCelebrity(*model.Celebrity)
	FetchCelebrity(string) *model.RawCelebrity

	CreateImageSet(*model.ImageSet)
	FetchImageSet(string) *model.RawImageSet

	CreateMedia(*model.Media)
	FetchMedia(string) *model.RawMedia

	CreateMovie(*model.Movie)
	FetchMovie(string) *model.RawMovie

	CreatePlaylist(*model.Playlist)
	FetchPlaylist(string) *model.RawPlaylist

	CreateSeason(*model.Season)
	FetchSeason(string) *model.RawSeason

	CreateSeries(*model.Series)
	FetchSeries(string) *model.RawSeries

	CreateSnippet(*model.Snippet)
	FetchSnippet(string) *model.RawSnippet

	CreateSong(*model.Song)
	FetchSong(string) *model.RawSong

	CreateTitle(*model.Title)
	FetchTitle(string) *model.RawTitle
}
