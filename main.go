package main

import (
	"encoding/json"
	"fmt"

	"github.com/lea-video/backend/storage"
	"github.com/lea-video/backend/types"
)

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func addVideo() *types.VideoMedia {
	title := types.Title{
		Title:     "Test Video",
		Language:  "en",
		PrevTitle: nil,
		NextTitle: nil,
	}
	video := types.Video{
		Title:     &title,
		Length:    128,
		AudioF:    make([]*types.AudioFormat, 0),
		VideoF:    make([]*types.VideoFormat, 0),
		SubtitleF: make([]*types.SubtitleFormat, 0),
	}
	mediaTitle := title
	media := types.Media{
		Title: &mediaTitle,
		Tags:  make([]*types.Tag, 0),
	}
	return &types.VideoMedia{
		Media: &media,
		Video: &video,
	}
}

func demoAddVideo() {
	// Create Video
	v := addVideo()
	fmt.Println("zero: " + v.String())

	// Build JSON from Video
	b, err := json.Marshal(v)
	panicOn(err)
	fmt.Println("first: " + string(b))

	// Modify original VIdeo to validate that the second log is not just a relog
	v.Title = nil

	// Load Video from JSON
	dat := types.VideoMedia{}
	err = json.Unmarshal(b, &dat)
	panicOn(err)
	// Build another JSON to log
	b, err = json.Marshal(dat)
	panicOn(err)
	fmt.Println("second: " + string(b))
}

func demoStorage() {
	s := storage.NewVideoStorage(&types.MySQL{}, 3)
	s.GetVideo("10")
	s.GetVideo("20")
	s.GetVideo("30")
	s.GetVideo("30")
	fmt.Println(s)
	s.GetVideo("40")
	s.GetVideo("50")
	s.GetVideo("40")
	fmt.Println(s)
}

func main() {
	demoStorage()
}
