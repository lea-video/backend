package types

type Format struct {
	Filename string
	Filesize uint
	Language string
}

type AudioFormat struct {
	*Format
	Bitrate  uint
	Encoding string
}

type VideoFormat struct {
	*Format
	Bitrate  uint
	Encoding string
}

type SubtitleFormat struct {
	*Format
}
