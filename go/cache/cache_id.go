package cache

import (
	"errors"
	"fmt"

	db "github.com/lea-video/backend/go/db"
	"github.com/lea-video/backend/go/model"
)

type StringSet map[string]interface{}

var storage StringSet

var empty interface{}

var ready bool

func InitCache(db db.DBInterface) {
	storage = make(StringSet)
	ready = false
	fetch(db)
}

func fetch(db db.DBInterface) {
	ids := db.FetchIDs()
	for i := 0; i < len(ids); i++ {
		WasCreated(ids[i])
	}
	ready = true
}

func IsReady() bool {
	return ready
}

func Exists(id string) bool {
	if !ready {
		panic(errors.New("cache not initialised"))
	}
	_, found := storage[id]
	return found
}

func IsCached(id string) bool {
	if !ready {
		panic(errors.New("cache not initialised"))
	}
	obj, found := storage[id]
	return found && obj != empty
}

func Cache(id string, obj interface{}) {
	if !ready {
		panic(errors.New("cache not initialised"))
	}
	storage[id] = obj
}

func WasCreated(id string) {
	if !ready {
		panic(errors.New("cache not initialised"))
	}
	storage[id] = empty
}

func Retrieve(id string) interface{} {
	return storage[id]
}

func WasDeleted(id string) {
	if !ready {
		panic(errors.New("cache not initialised"))
	}
	delete(storage, id)
}

func RetrieveTitel(id string) (t *model.Title, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Title), nil
}

func RetrieveTitleItem(id string) (t *model.TitleItem, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.TitleItem), nil
}

func RetrieveSong(id string) (t *model.Song, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Song), nil
}

func RetrieveSongCover(id string) (t *model.SongCover, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.SongCover), nil
}

func RetrieveTag(id string) (t *model.Tag, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Tag), nil
}

func RetrieveMedia(id string) (t *model.Media, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Media), nil
}

func RetrieveTrack(id string) (t *model.Track, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Track), nil
}

func RetrieveMovie(id string) (t *model.Movie, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Movie), nil
}

func RetrieveSeries(id string) (t *model.Series, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Series), nil
}

func RetrieveSeason(id string) (t *model.Season, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Season), nil
}

func RetrieveCelebrity(id string) (t *model.Celebrity, err error) {
	obj, found := storage[id]
	if !found || obj == empty {
		return nil, errors.New("unknown")
	}
	defer func() {
		if recover() != nil {
			t = nil
			err = errors.New("cast failed")
		}
	}()
	return obj.(*model.Celebrity), nil
}

func Test() {
	fmt.Println("start", storage)
	// set some vars
	WasCreated("1234")
	Cache("2345", model.NewTitle("Test 1"))
	Cache("3456", model.NewCelebrity("Test 2"))
	fmt.Println(storage)
	// check bool methods
	fmt.Println(Exists("1234"), IsCached("1234"))
	fmt.Println(Exists("2345"), IsCached("2345"))
	fmt.Println(Exists("3456"), IsCached("3456"))
	// check retrieval
	fmt.Println(RetrieveTitel("01234"))
	fmt.Println(RetrieveTitel("1234"))
	fmt.Println(RetrieveTitel("2345"))
	fmt.Println(RetrieveTitel("3456"))
	fmt.Println("end")
}
