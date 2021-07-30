package types

type Celebrity struct {
	Name string
	Tags []*Tag
}

type Playlist struct {
	Title *Title
	Items []*PlaylistItem
}

type PlaylistItem struct {
	Idx       uint
	Item      *Media
	AddedAt   uint
	RemovedAt uint
}
