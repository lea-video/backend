package types

type Title struct {
	Title     string
	Language  string
	PrevTitle *Title
	NextTitle *Title
}

type Tag struct {
	Title *Title
}
