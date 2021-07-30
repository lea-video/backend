package main

import (
	"fmt"

	"github.com/lea-video/backend/types"
)

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

func main() {
	v := addVideo()
	fmt.Println(v)
}
