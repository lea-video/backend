package storage

import (
	"github.com/lea-video/backend/types"
)

type VideoStorage struct {
	db       types.DBInterface
	cache    []*types.Video
	idMap    map[string]int
	revIdMap map[int]string
	pos      int
}

func NewVideoStorage(db types.DBInterface, size int) *VideoStorage {
	return &VideoStorage{
		db:       db,
		cache:    make([]*types.Video, size),
		idMap:    make(map[string]int),
		revIdMap: make(map[int]string),
	}
}

func (vs *VideoStorage) GetVideo(id string) *types.Video {
	cacheIdx, isCached := vs.idMap[id]
	if isCached {
		return vs.cache[cacheIdx]
	}
	// fetch video
	video := vs.db.GetVideo(id)
	// calculate and save new vs.pos
	cacheIdx = vs.pos
	vs.pos = (vs.pos + 1) % len(vs.cache)
	// fetch id of prev cached item and delete
	prevCached, hadPrev := vs.revIdMap[cacheIdx]
	if hadPrev {
		delete(vs.idMap, prevCached)
	}
	// cache new item information
	vs.idMap[id] = cacheIdx
	vs.revIdMap[cacheIdx] = id
	vs.cache[cacheIdx] = video
	// done
	return video
}
