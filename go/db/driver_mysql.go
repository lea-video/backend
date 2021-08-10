package db

import (
	"fmt"

	model "github.com/lea-video/backend/go/model"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	user         string
	password     string
	dbname       string
	ressourceStr string
	db           *sql.DB
}

func NewMySQL(user, password, dbname string) *mysql {
	res := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbname)
	db, err := sql.Open("mysql", res)
	panicOn(err)
	return &mysql{
		user:         user,
		password:     password,
		dbname:       dbname,
		ressourceStr: res,
		db:           db,
	}
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func (m *mysql) FetchIDs() []string {
	return make([]string, 0)
}

func (m *mysql) CreateAlbum(*model.Album) {}
func (m *mysql) FetchAlbum(id string) *model.RawAlbum {
	// TODO: test
	rows, err := m.db.Query("SELECT * FROM Album WHERE ID=?", id)
	panicOn(err)
	dest := &model.RawAlbum{}
	rows.Scan(dest)
	return dest
}

func (m *mysql) CreateCelebrity(*model.Celebrity) {}
func (m *mysql) FetchCelebrity(string) *model.RawCelebrity {
	return nil
}

func (m *mysql) CreateImageSet(*model.ImageSet) {}
func (m *mysql) FetchImageSet(string) *model.RawImageSet {
	return nil
}

func (m *mysql) CreateMedia(*model.Media) {}
func (m *mysql) FetchMedia(string) *model.RawMedia {
	return nil
}

func (m *mysql) CreateMovie(*model.Movie) {}
func (m *mysql) FetchMovie(string) *model.RawMovie {
	return nil
}

func (m *mysql) CreatePlaylist(*model.Playlist) {}
func (m *mysql) FetchPlaylist(string) *model.RawPlaylist {
	return nil
}

func (m *mysql) CreateSeason(*model.Season) {}
func (m *mysql) FetchSeason(string) *model.RawSeason {
	return nil
}

func (m *mysql) CreateSeries(*model.Series) {}
func (m *mysql) FetchSeries(string) *model.RawSeries {
	return nil
}

func (m *mysql) CreateSnippet(*model.Snippet) {}
func (m *mysql) FetchSnippet(string) *model.RawSnippet {
	return nil
}

func (m *mysql) CreateSong(*model.Song) {}
func (m *mysql) FetchSong(string) *model.RawSong {
	return nil
}

func (m *mysql) CreateTitle(*model.Title) {}
func (m *mysql) FetchTitle(string) *model.RawTitle {
	return nil
}
