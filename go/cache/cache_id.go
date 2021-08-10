package cache

import (
	"errors"

	db "github.com/lea-video/backend/go/db"
)

type StringSet map[string]struct{}

var storage StringSet

var empty struct{}

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

func WasCreated(id string) {
	if !ready {
		panic(errors.New("cache not initialised"))
	}
	storage[id] = empty
}

func WasDeleted(id string) {
	if !ready {
		panic(errors.New("cache not initialised"))
	}
	delete(storage, id)
}
