package types

type Media struct {
	Title *Title
	Video []*Video
	Tags  []*Tag
}

type VideoMedia struct {
	Video
}

type MovieMedia struct {
	Cuts     []*Video
	Actors   []*Celebrity
	Director []*Celebrity
}

type MusikMedia struct {
	Variants   []*Video
	Artists    []*Celebrity
	CleanAudio *Media
}
