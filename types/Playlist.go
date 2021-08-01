package types

type Celebrity struct {
	Name        string `json:"name"`
	ProfileLink string `json:"profileLink"`
	Tags        []*Tag `json:"tags"`
}

type Playlist struct {
	Title *Title          `json:"title"`
	Items []*PlaylistItem `json:"items"`
}

type PlaylistItem struct {
	Idx       uint   `json:"idx"`
	Item      *Media `json:"item"`
	AddedAt   uint   `json:"addedAt"`
	RemovedAt uint   `json:"removedAt"`
}
