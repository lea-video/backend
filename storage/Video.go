package storage

import (
	"fmt"

	"github.com/lea-video/backend/types"
)

type VideoStorage struct {
	db    types.DBInterface
	cache []*types.Video
	items map[string]int
	pos   int
}

func NewVideoStorage(db types.DBInterface, size int) *VideoStorage {
	return &VideoStorage{
		db:    db,
		cache: make([]*types.Video, size),
		items: make(map[string]int),
	}
}

func (vs *VideoStorage) GetVideo(id string) *types.Video {
	fmt.Printf("GetVideo(%s)\n", id)
	cacheIdx := vs.items[id]
	if cacheIdx != 0 {
		fmt.Println("cache hit")
		return vs.cache[cacheIdx]
	}
	video := vs.db.GetVideo(id)
	cacheIdx = (vs.pos + 1) % len(vs.cache)
	fmt.Printf("fetching from db and saving in %d\n", cacheIdx)
	delete(vs.items, vs.cache[cacheIdx].ID)
	vs.items[id] = cacheIdx
	vs.cache[cacheIdx] = video
	return video
}
