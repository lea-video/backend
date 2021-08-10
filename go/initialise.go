package leav

import (
	cache "github.com/lea-video/backend/go/cache"
	db "github.com/lea-video/backend/go/db"
	util "github.com/lea-video/backend/go/utility"
)

func Initialise() {
	util.Init()
	cache.InitCache(db.NewMySQL("testuser", "testpw", "testdb"))
	// wait for cache to be initialised
	for !cache.IsReady() {
	}
}
