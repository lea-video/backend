package types

type Media struct {
	Title *Title
	Tags  []*Tag
}

type VideoMedia struct {
	*Media
	Video *Video
}

type MovieMedia struct {
	*Media
	Cuts     []*Video
	Actors   []*Celebrity
	Director []*Celebrity
}

type MusikMedia struct {
	*Media
	Variants   []*Video
	Artists    []*Celebrity
	CleanAudio *VideoMedia
	BestVideo  *VideoMedia
}
