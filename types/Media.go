package types

import "fmt"

type Media struct {
	Title *Title `json:"title"`
	Tags  []*Tag `json:"tags"`
}

type VideoMedia struct {
	*Media
	Video *Video `json:"video"`
}

func (t *VideoMedia) String() string {
	return fmt.Sprintf("VideoMedia with title '%s' for Video '%s'", t.Title, t.Video)
}

type MovieMedia struct {
	*Media
	Cuts      []*Video     `json:"videos"`
	Actors    []*Celebrity `json:"actors"`
	Directors []*Celebrity `json:"directors"`
}

type MusikMedia struct {
	*Media
	Variants   []*Video     `json:"variants"`
	Artists    []*Celebrity `json:"artists"`
	CleanAudio *VideoMedia  `json:"cleanAudio"`
	BestVideo  *VideoMedia  `json:"bestVideo"`
}
